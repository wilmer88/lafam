package main

import (
	"os"
	"gorm.io/gorm"
	"lafam/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "gorm-test/controllers"
	"net/http"
)

func main() {

	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	_ = r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	config.AllowOrigins = []string{"https://localhost:4200/"}
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("localhost:4200/ *", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	return r
}
