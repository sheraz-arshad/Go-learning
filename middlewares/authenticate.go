package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(f gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		if session.Get("signed_in") != true {
			context.JSON(http.StatusForbidden, gin.H{
				"error": "not allowed",
			})
			return
		}
		f(context)
	}
}
