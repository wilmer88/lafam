package main

import (
	"os"

	"lafam/controllers"

	"github.com/gin-gonic/gin"

	// "gorm-test/controllers"
	"net/http"
)

func main() {

	let port = process.env.PORT;,

if (port == null || port == "") {
  port = 8000;
}


	r := setupRouter()
	_ = r.Run(":"+port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	// r.POST("/lafamily", userRepo.CreateUser)
	// r.GET("/lafamily", userRepo.GetUsers)
	// r.GET("/lafamily/:id", userRepo.GetUser)
	// r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	return r
}