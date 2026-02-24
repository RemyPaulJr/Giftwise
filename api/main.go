package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Struct for receiving email and password input from user
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User struct for passing to database
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

func main() {
	// Creates new chi router instance that will handle all incoming requests and route them to the right handler function
	r := chi.NewRouter()
	// Applies middleware to every route on this router. auto logs every incoming request
	r.Use(middleware.Logger)

	r.Post("/auth/register", registerUser)
}

func registerUser(w http.ResponseWriter, r *http.Request) {

}
