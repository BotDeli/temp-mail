package cookie

import "github.com/gin-gonic/gin"

func GetUUIDFromCookie(ctx *gin.Context) (string, error) {
	return ctx.Cookie("uuid")
}

func SetUUIDToCookie(ctx *gin.Context, uuid string) {
	ctx.SetCookie("uuid", uuid, 30*60, "/", "", false, true)
}
