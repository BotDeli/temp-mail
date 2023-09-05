package handlers

import (
	"github.com/gin-gonic/gin"
	"temp-mail/internal/storage/connections"
	"temp-mail/internal/storage/mails"
)

func InitHandlers(router *gin.Engine, sm mails.StorageMails, sc connections.StorageConnections) {
	router.GET("/", loadMainPage)
	router.GET("/get_mail", handlerGetMail(sm))
	router.GET("/getMessages", handleWBGetMessages(sm, sc))
	router.GET("/message", handlerGetMessage())
}
