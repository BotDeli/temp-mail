package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"temp-mail/internal/server/handlers"
	"temp-mail/internal/storage/connections"
	"temp-mail/internal/storage/mails"
)

func StartServer(sm mails.StorageMails, sc connections.StorageConnections) error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	loadFiles(router)
	address := os.Getenv("server_address") + ":" + os.Getenv("server_port")
	handlers.InitHandlers(router, sm, sc)
	return router.Run(address)
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/images", "./images")
	router.LoadHTMLGlob("templates/*")
}
