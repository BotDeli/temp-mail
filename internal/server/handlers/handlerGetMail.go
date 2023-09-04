package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"temp-mail/internal/server/cookie"
	"temp-mail/internal/storage/mails"
)

func handlerGetMail(s mails.StorageMails) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mail, err := getMail(ctx, s)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(200, gin.H{"mail": mail})
	}
}

func getMail(ctx *gin.Context, s mails.StorageMails) (string, error) {
	uuid, err := cookie.GetUUIDFromCookie(ctx)
	if err == nil {
		if mail, err := s.GetMail(uuid); err == nil {
			return mail, nil
		}
	}

	uuid, mail, err := s.NewMail()
	if err != nil {
		return "", err
	}

	cookie.SetUUIDToCookie(ctx, uuid)
	return mail, nil
}
