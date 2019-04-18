package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type SocketClient struct {
	Reader chan string
	Sender chan string
	socket *websocket.Conn
}

var UhomeSocket *SocketClient

func (s *SocketClient) send() {
	defer func() {
		s.socket.Close()
	}()

	for {
		select {
		case msg := <-s.Sender:
			err := s.socket.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println(" **  socket send err")
			}
		}
	}
}

func (s *SocketClient) read() {
	defer func() {
		s.socket.Close()
	}()

	var receiveMSG struct {
		Event string
		Type  string
		Data  string
	}

	for {
		err := s.socket.ReadJSON(&receiveMSG)
		if err != nil {
			resp, _ := json.Marshal(map[string]string{
				"event": "error",
				"type":  "server",
				"data":  "receive msg err",
			})
			s.Sender <- string(resp)
		}

		if receiveMSG.Event == "ping" {
			resp, _ := json.Marshal(map[string]string{"event": "pong"})
			s.Sender <- string(resp)
		}
	}
}
