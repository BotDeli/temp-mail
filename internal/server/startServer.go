package server

import (
	"github.com/gin-gonic/gin"
	"temp-mail/internal/server/handlers"
)

func StartServer() error {
	router := gin.Default()
	loadFiles(router)
	handlers.InitHandlers(router)
	return router.Run(":8080")
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/images", "./images")
	router.LoadHTMLGlob("templates/*")
}
