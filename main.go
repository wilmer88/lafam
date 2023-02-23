package main

import (
	// "embed"
	"log"

	"os"

	// "crypto/tls"
	// "net/http"

	"github.com/gin-contrib/cors"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	// "github.com/wilmer88/lafam/controllers"
)


func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
    //     w.Write([]byte("Hello, TLS!"))
    // })

	// http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	// http.Handle("http://localhost:8080", http.FileServer(http.Dir("./client/dist/client")))
	// http.Handle("https://mifamily-app.herokuapp.com", http.FileServer(http.Dir("./public/client/dist")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	// r := setupRouter()
	
	// r.Static( "/public", "./client/dist" )


	r := setupRouter()
	// r.StaticFileFS( "http:localhost:8080", "./public/client/dist/index.html" )
	r.Static( "/", "./public/dist/client/index.html")
	// r.GET("/lafamily", func (c *gin.Context)  {
		// c.File(http.StatusOK, "./public/dist/client/index.html")	
	// 	c.String(http.StatusOK, "working")	

	// })
	
    log.Fatal( r.Run(":"+port ))
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

func setupRouter()*gin.Engine {
	
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com"}
	// config.AllowOrigins = []string{"http://localhost:4200"}
	r.Use(cors.New(config))

	// userRepo := controllers.New()
	// r.POST("/lafamily", userRepo.CreateUser)
	// r.Static(userRepo.GetUser,"/lafamily/:id" )
	// r.PUT("/lafamily/:id", userRepo.UpdateUser)
	// r.DELETE("/lafamily/:id", userRepo.DeleteUser)
	// r.Use(static.Serve("https://mifamily-app.herokuapp.com/lafamily", static.LocalFile("./public/client",true))) 
	
	return r
}

