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

type ReceivedMSG struct {
	Event string
	Type  string
	Data  string
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
				log.Println(" **  socket send err: ", err)
			}
		}
	}
}

func (s *SocketClient) read() {
	defer func() {
		s.socket.Close()
	}()

	var resp []byte

	for {
		_, received, err := s.socket.ReadMessage()
		if err != nil {
			log.Println(" **  socket receive msg err: ", err)
			s.socket.Close()
			break
		}

		var parsedMSG ReceivedMSG
		err = json.Unmarshal(received, &parsedMSG)
		if err != nil {
			resp, _ = json.Marshal(map[string]string{
				"event": "error",
				"type":  "serve",
				"data":  "parse received msg err",
			})
			s.Sender <- string(resp)
			continue
		}

		if parsedMSG.Event == "ping" {
			resp, _ := json.Marshal(map[string]string{"event": "pong"})
			s.Sender <- string(resp)
		}
	}
}
