package main

import (
    // "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "path/filepath"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/wilmer88/lafam/controllers"
)

func main() {

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    r := setupRouter()

    // Create a reverse proxy for requests to the Angular Compiler CLI WebSocket server
    // wsTarget, _ := url.Parse("https://mifamily-app.herokuapp.com")
    wsTarget, _ := url.Parse("http://localhost:5000")
    wsProxy := httputil.NewSingleHostReverseProxy(wsTarget)

    // Add a handler for requests to /ngc-cli-ws that passes the request through to the WebSocket server
    r.NoRoute(func(c *gin.Context) {
        if c.Request.URL.Path == "/ngc-cli-ws" {
            wsProxy.ServeHTTP(c.Writer, c.Request)
        }
    })

    // Start the server
    _ = r.Run(":" + port)
}

func setupRouter() *gin.Engine {
    r := gin.Default()

    // Set up CORS
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com, http://localhost:5000"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    r.Use(cors.New(config))

    // Set up static file serving
    distDir := filepath.Join(".", "public/dist/public")
    r.StaticFile("/index.html", filepath.Join(distDir, "index.html"))

    // Set up dynamic routes
    userRepo := controllers.New()
    r.POST("/lafamily", userRepo.CreateUser)
    r.GET("/", userRepo.GetUsers)
    r.GET("/lafamily/user/:id", userRepo.GetUser)
    r.PUT("/lafamily/user/:id", userRepo.UpdateUser)
    r.DELETE("/lafamily/user/:id", userRepo.DeleteUser)

    return r
}