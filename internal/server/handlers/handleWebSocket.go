package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var socket = websocket.Upgrader{
	HandshakeTimeout: 30 * time.Minute,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		msg, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg, string(p))

		response := []byte("Test")
		err = conn.WriteMessage(msg, response)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
