package handlers

import (
	"github.com/gin-gonic/gin"
)

type PageData struct {
	Sender  string
	Date    string
	Subject string
	Content string
}

func handlerGetMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loadViewPage(ctx, "testSender", "testDate", "testSubject", "testContent")
	}
}

func loadViewPage(ctx *gin.Context, sender, date, subject, content string) {
	pd := PageData{
		Sender:  sender,
		Date:    date,
		Subject: subject,
		Content: content,
	}
	ctx.HTML(200, "view.html", pd)
}
