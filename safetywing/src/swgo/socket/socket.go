package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender    string
	Recipient string
	Content   string
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan map[string]string
}

type ClientManager struct {
	broadcast  chan map[string]string
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
}

var Manager = ClientManager{
	broadcast:  make(chan map[string]string),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

// Start ...
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			log.Println("a new socket has connected")
			mapMess := map[string]string{"id": conn.id, "type": "open", "message": "opened ws"}
			conn.send <- mapMess
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				log.Println("a socket has disconnected")
				// 关闭了,,,不应该发送消息了
				// mapMess := map[string]string{"id": conn.id, "type": "close", "message": "ws closed"}
				// conn.send <- mapMess
			}
		case message := <-manager.broadcast:
			manager.Send(message)
		}
	}
}

// Send ...
// message map should be ["id":"", "type":"", "message":""]
func (manager *ClientManager) Send(message map[string]string) {
	for conn := range manager.clients {
		conn.send <- message
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			jsonMess, _ := json.Marshal(message)
			c.socket.WriteMessage(websocket.TextMessage, jsonMess)
		}
	}
}

func (c *Client) read() {
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.socket.Close()
			break
		}
		// broadcastMess, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		broadCastMess := map[string]string{"id": c.id, "type": "message", "message": string(message)}
		Manager.broadcast <- broadCastMess
	}
}

// Upgrade2WS ...
func Upgrade2WS(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade err", 500)
		return
	}

	newuuid, _ := uuid.NewV4()
	newid := newuuid.String()
	client := &Client{id: newid, socket: conn, send: make(chan map[string]string)}
	Manager.register <- client

	go client.read()
	go client.write()
}

// MockPostSome ...
func MockPostSome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("parse form err")
		http.Error(w, "parse form err", 400)
		return
	}

	mapMessage := map[string]string{"id": "", "type": "message", "message": r.Form.Get("message")}
	Manager.Send(mapMessage)
	w.Write([]byte("post ok"))
	return
}
