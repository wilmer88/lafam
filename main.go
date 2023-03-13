package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/controllers"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	_ = r.Run(":" + port)

}

func setupRouter() *gin.Engine {
	// gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Set up CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com"}
	// config.AllowOrigins = []string{"https://localhost:8080/lafamily"}
	// config.AllowOrigins = []string{"http://localhost:4200"}

	r.Use(cors.New(config))

	// Set up static file serving
	static := r.Group("/static")
	static.Static("/", "./client/dist/index.html")
	// static.Static("/", "./public/dist/client/index.html")


	// Set up dynamic routes
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/user/:id", userRepo.GetUser)
	r.PUT("/lafamily/user/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/user/:id", userRepo.DeleteUser)

	return r
}
