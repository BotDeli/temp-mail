package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"temp-mail/internal/api/gMessages"
)

func handlerGetMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mail := ctx.Query("mail")
		id := ctx.Query("id")
		if mail == "" || id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		msg, err := gMessages.ReadMessage(mail, id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		loadViewPage(ctx, msg.From, msg.Date, msg.Subject, repairHTML(msg.Body))
	}
}

func repairHTML(str string) string {
	str = strings.ReplaceAll(str, "<HTML>", "")
	str = strings.ReplaceAll(str, "</HTML>", "")
	str = strings.ReplaceAll(str, "<BODY>", "")
	str = strings.ReplaceAll(str, "</BODY>", "")
	str = strings.ReplaceAll(str, "<div>", "")
	str = strings.ReplaceAll(str, "</div>", "")
	return str
}
