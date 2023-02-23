package main

import (
	"os"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := chi.NewRouter()
	r.Get("http://localhost:4200/family", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("im working"))
	})

	http.ListenAndServe(":"+port, r)

}
