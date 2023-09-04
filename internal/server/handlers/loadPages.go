package handlers

import "github.com/gin-gonic/gin"

func loadMainPage(ctx *gin.Context) {
	ctx.HTML(200, "main.html", nil)
}
