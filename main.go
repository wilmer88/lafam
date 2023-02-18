package main

import (
	"os"
	"lafam/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "gorm-test/controllers"
	"net/http"
	"github.com/memcachier/mc"
	"github.com/gin-contrib/cache"
  "github.com/gin-contrib/cache/persistence"
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
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com, https://localhost:4200"}



	r.GET("/api/lafamily", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
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