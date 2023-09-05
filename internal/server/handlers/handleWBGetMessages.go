package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"temp-mail/internal/api/gMessages"
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

		mail, err := sm.GetMail(uuid)

		var messages []gMessages.InputMessage

		for sc.IsLifeConnection(uuid) {
			//msg, p, _ := conn.ReadMessage()
			//fmt.Println(msg, string(p))

			messages, err = gMessages.GetMessagesFromMail(mail)
			if err == nil {
				err = conn.WriteJSON(messages)
			}
			time.Sleep(5 * time.Second)
		}
		closeConnection(conn)
	}
}

func closeConnection(conn *websocket.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}
