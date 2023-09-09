package handlers

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func loadMainPage(ctx *gin.Context) {
	ctx.HTML(200, "main.html", nil)
}

func loadViewPage(ctx *gin.Context, sender, date, subject, content string) {
	ctx.HTML(200, "view.html", gin.H{
		"Sender":  sender,
		"Date":    date,
		"Subject": subject,
		"Content": template.HTML(content),
	})
}
