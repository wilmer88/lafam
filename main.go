package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/controllers"

	// "gorm-test/controllers"
	"net/http"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

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