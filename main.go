package main

import (
	"os"
	"github.com/wilmer88/lafam/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "gorm-test/controllers"
	"net/http"
)

func main() {

	port := os.Getenv("Port")
	if port == "" {
		port = ":8080"
	}

	r := setupRouter()
	_ = r.Run(":8080" )
}

func setupRouter() *gin.Engine {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://mifamily-app.herokuapp.com"}
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.Use(cors.New(config))
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)
	

	return r
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}