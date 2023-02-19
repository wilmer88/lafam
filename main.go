package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gin-contrib/cors"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/memcachier/mc"
	"github.com/wilmer88/lafam/controllers"
)

func main() {
	username := os.Getenv("MEMCACHIER_USERNAME")
	password := os.Getenv("MEMCACHIER_PASSWORD")
	servers := os.Getenv("MEMCACHIER_SERVERS")

	mcClient := mc.NewMC(servers, username, password)
	defer mcClient.Quit()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://localhost:4200"}
	r.Use(cors.New(config))
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("/public")
		} else {
			c.File("/dist" + path.Join(dir, file))
		}
	})
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	_ = r.Run(":" + port)

}
