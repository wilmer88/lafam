package main

import (
	"os"
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
	// r.GET("ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, "pong")
	// })
	config.AllowOrigins = []string{"https://localhost:4200/"}
	r.Use(cors.New(config))
	userRepo := controllers.New()

	r.GET("/", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)
	r.POST("/lafamily", userRepo.CreateUser)

	return r
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}