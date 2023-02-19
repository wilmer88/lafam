package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/wilmer88/lafam/controllers"
	"github.com/memcachier/mc"
)

func main() {
	username := os.Getenv("MEMCACHIER_USERNAME")
	password := os.Getenv("MEMCACHIER_PASSWORD")
	servers := os.Getenv("MEMCACHIER_SERVERS")
  
	mcClient := mc.NewMC(servers, username, password)
	defer mcClient.Quit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	

	r := setupRouter()

	r.Use(gin.Static("/public"))

	_ = r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com"}
	r.Use(cors.New(config))

	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	return r
}