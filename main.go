package main

import (
	"embed"

	"os"

	// "crypto/tls"
	"net/http"

	"github.com/gin-contrib/cors"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/controllers"
)




//go:embed public/*
var f embed.FS
func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
    //     w.Write([]byte("Hello, TLS!"))
    // })

	// http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	// http.Handle("http://localhost:4200/lafamily", http.FileServer(http.Dir("./client/dist/client")))
	// http.Handle("https://mifamily-app.herokuapp.com", http.FileServer(http.Dir("./public/client/dist")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	// r := setupRouter()
	
	// r.Static( "/public", "./client/dist" )


	r := gin.Default()
	r.StaticFS( "fs", http.FileSystem(http.FS(f)))
	r.GET("/family", func (c *gin.Context)  {
		c.File("./public/dist/client/index.html")	
	})
	_ = r.Run(":"+port )
	// tlsConfig := &tls.Config{
	// 	MinVersion: tls.VersionTLS12,
	// }
	
	// server := &http.Server{
	// 	Addr:      ":8080",
	// 	TLSConfig: tlsConfig,
	// 	Handler:   r,
	// }
	
	// if err := server.ListenAndServeTLS("path/to/cert.pem", "path/to/key.pem"); err != nil {
	// 	panic(err)
	// }
}

func setupRouter() *gin.Engine {
	
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com"}
	// config.AllowOrigins = []string{"http://localhost:4200"}
	r.Use(cors.New(config))
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)
	// r.Use(static.Serve("https://mifamily-app.herokuapp.com/lafamily", static.LocalFile("./public/client",true))) 
	
	return r
}

