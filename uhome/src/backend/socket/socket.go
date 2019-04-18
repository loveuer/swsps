package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

// Init generate a websocket
func Init(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "socket -> Init: upgrade to websocket err", 500)
		return
	}

	UhomeSocket := &SocketClient{
		Reader: make(chan string, 1),
		Sender: make(chan string, 1),
		socket: conn,
	}

	go UhomeSocket.send()
	go UhomeSocket.read()

	return
}
