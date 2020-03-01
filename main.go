package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/middlewares"
	"learning/models"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("1234"))
	router.Use(
		sessions.Sessions(
			"my_session",
			store,
		),
	)

	db := models.SetupModels()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.POST("/signup", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)

	authenticate := middlewares.Authenticate

	router.GET("/users", authenticate(controllers.GetUsers))
	router.GET("/users/:id", authenticate(controllers.GetUserById))
	router.PUT("/users/:id", authenticate(controllers.UpdateUser))
	router.DELETE("/users/:id", authenticate(controllers.DeleteUser))

	router.GET("/book/:id", authenticate(controllers.FindBook))
	router.GET("/books", authenticate(controllers.FindBooks))
	router.POST("/book", authenticate(controllers.CreateBook))
	router.PUT("/book/:id", authenticate(controllers.UpdateBook))
	router.DELETE("/book/:id", authenticate(controllers.DeleteBook))
	router.Run()
}
