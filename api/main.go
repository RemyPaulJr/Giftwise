package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// User struct for passing to database
type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Password_Hash string `json:"password_hash"`
	Created_At    time.Time
}

func main() {
	// Creates new chi router instance that will handle all incoming requests and route them to the right handler function
	r := chi.NewRouter()
	// Applies middleware to every route on this router. auto logs every incoming request
	r.Use(middleware.Logger)
	// Connect to database and create app instance
	app, err := startDB()
	if err != nil {
		log.Print("Failed to start app server: ", err)
		return
	}

	r.Post("/auth/register", app.registerUser)
	r.Post("/auth/login", app.loginUser)

	log.Print("Starting server on :8080")
	http.ListenAndServe(":8080", r)

}
