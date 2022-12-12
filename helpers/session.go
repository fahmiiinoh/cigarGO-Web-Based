package helpers

import (
	"github.com/gin-gonic/gin"
)

func SessionSet(ctx *gin.Context, userID uint64) {
	session := sessions.Default(ctx)
	var idInterface interface{} = &userID
	session.Set("id", idInterface)
	session.Save()

}

func SessionGet(ctx *gin.Context) uint64 {
	session := sessions.Default(ctx)
	return session.Get("id").(uint64)
}

func SessionClear(ctx *gin.Context) {

	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}
