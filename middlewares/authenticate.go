package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Authenticate(f gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Query("token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			context.Set("user", claims["name"])
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		f(context)
	}
}
