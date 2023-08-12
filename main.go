package main

import (
	"os"
	// "github.com/gorilla/mux"
	"github.com/gin-contrib/cors"
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/controllers"
)


// func main() {
// 	r := mux.NewRouter()
  
// 	r.HandleFunc("/hello-world", helloWorld)
  
// 	http.Handle("/", r)
  
// 	srv := &http.Server{
// 	  Handler: r,
// 	  Addr:    ":" + os.Getenv("PORT"),
// 	}
  
// 	log.Fatal(srv.ListenAndServe())
//   }
  
//   func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	var data = struct {
// 	  Title string `json:"title"`
// 	}{
// 	  Title: "Golang + Angular Starter Kit",
// 	}
  
// 	jsonBytes, err := utils.StructToJSON(data); if err != nil {
// 	  fmt.Print(err)
// 	}
  
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonBytes)
// 	return
//   }

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := setupRouter()
	_ = r.Run("0.0.0.0:" + port)

}
func setupRouter() *gin.Engine {
	// gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Set up CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mifamily-app.herokuapp.com", "http://localhost:4200"}
	// config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// Set up static file serving
	// static := r.Group("/static")
	// static.Static("/public", "./client/dist")
	// r.StaticFile("/public", "./client/dist/index.html")
	r.Static("/public", "./client/dist")



	// Set up dynamic routes
	userRepo := controllers.New()
	r.POST("/lafamily/post", userRepo.CreateUser)
	r.GET("/api/lafam", userRepo.GetUsers)
	// r.GET("https://mifamily-app.herokuapp.com", userRepo.GetUsers)
	r.GET("/lamily/user/:id", userRepo.GetUser)
	r.PUT("/lafamily/user/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/user/:id", userRepo.DeleteUser)

	return r
}
