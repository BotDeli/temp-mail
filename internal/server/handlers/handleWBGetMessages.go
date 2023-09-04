package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"temp-mail/internal/server/cookie"
	"temp-mail/internal/storage/connections"
	"temp-mail/internal/storage/mails"
	"time"
)

var socket = websocket.Upgrader{
	HandshakeTimeout: 30 * time.Minute,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

func handleWBGetMessages(sm mails.StorageMails, sc connections.StorageConnections) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, err := cookie.GetUUIDFromCookie(ctx)
		if err != nil {
			return
		}

		for sc.IsLifeConnection(uuid) {
			sc.StopLifeConnection(uuid)
			time.Sleep(6 * time.Second)
		}

		conn, err := socket.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}
		sc.StartLifeConnection(uuid)
		defer closeConnection(conn)

		mail, err := sm.GetMail(uuid)

		for sc.IsLifeConnection(uuid) {
			//msg, p, err := conn.ReadMessage()
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			//fmt.Println(msg, string(p))
			//
			//response := []byte("Test")
			//err = conn.WriteMessage(msg, response)
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			fmt.Println(mail)
			time.Sleep(5 * time.Second)
		}
	}
}

func closeConnection(conn *websocket.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}
