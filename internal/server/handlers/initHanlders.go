package handlers

import "github.com/gin-gonic/gin"

func InitHandlers(router *gin.Engine) {
	router.GET("/", loadMainPage)
}
