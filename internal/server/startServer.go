package server

import (
	"github.com/gin-gonic/gin"
	"temp-mail/internal/server/handlers"
	"temp-mail/internal/storage/connections"
	"temp-mail/internal/storage/mails"
)

func StartServer(sm mails.StorageMails, sc connections.StorageConnections) error {
	router := gin.Default()
	loadFiles(router)
	handlers.InitHandlers(router, sm, sc)
	return router.Run(":8080")
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/images", "./images")
	router.LoadHTMLGlob("templates/*")
}
