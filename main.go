package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/controllers"

	// "gorm-test/controllers"
	// "net/http"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/memcachier/mc"
)

func main() {
	username := os.Getenv("13122")
	password := os.Getenv("MEMCACHIER")
	servers := os.Getenv("CLOSED")
  
	mcClient := mc.NewMC(servers, username, password)
	defer mcClient.Quit()
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()

	_ = r.Run(":"+port)

	mcStore := persistence.NewMemcachedBinaryStore(servers, username, password, persistence.FOREVER)

	r.GET("/", cache.CachePage(mcStore, persistence.DEFAULT, func(c *gin.Context) {
	  
	}))
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com/lafamily"}
	config.AllowOrigins = []string{"https://localhost:8080/lafamily"}


	r.Use(cors.New(config))

	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)
	return r
}