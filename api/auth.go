package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// Struct for receiving email and password input from user
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse struct to send response back to client in JSON
type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	var email string
	// Var for setting the computational cost for hashing algo
	const cost = 12
	var request RegisterRequest
	var response RegisterResponse

	// Read the request from user (r) (email and password) and store in var decoder
	decoder := json.NewDecoder(r.Body)
	// Take what we read from decoder and store the decoded JSON in our RegisterRequest struct
	err := decoder.Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Validate Email
	if _, err = mail.ParseAddress(request.Email); err != nil {
		checkError(w, http.StatusBadRequest, "Invalid email address")
		return
	}
	// Create response email
	response.Email = request.Email

	// Validate Password for character length 8+
	if len(request.Password) < 8 {
		checkError(w, http.StatusBadRequest, "Password does not meet length requirement 8 characters")
		return
	}

	// 72 bytes is the hard limit for bcrypt. bycrypt's blowfish cipher algo gurantee security for every byte up until 72, anything after is ignored.
	// if two users enter 100 char passwords and they are the same up until the 72nd byte they could share the same hash.
	if len(request.Password) > 72 {
		checkError(w, http.StatusBadRequest, "Password too long. Please try a shorter password and try again.")
	}

	// Check for email in db
	emailCheck := a.db.QueryRow(r.Context(), "SELECT email FROM users WHERE email= $1", request.Email)
	if err = emailCheck.Scan(&email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			// Hash Password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), cost)
			if err != nil {
				checkError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
				log.Print("Error hashing password: ", err)
				return
			}

			// Generate Random UUID and Create RegisterRepsonse ID
			id := uuid.New().String()
			response.ID = id

			// Insert UUID, email, and hashpassword into DB
			commandTag, err := a.db.Exec(r.Context(), "INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3)", id, request.Email, string(hashedPassword))
			if err != nil {
				checkError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
				log.Print("Error storing user to database: ", err)
				return
			}
			checkError(w, http.StatusAccepted, "Successfully created user.")
			log.Printf("%d row affected.", commandTag.RowsAffected())
			return

		} else {
			checkError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
			log.Print("Error querying users table: ", err)
		}
	} else {
		checkError(w, http.StatusBadRequest, "User already exists.")
		log.Print("User already exists: ", err)
		return
	}

	writeJSON(w, http.StatusCreated, response)
}

func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {

	var request RegisterRequest
	var password string

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Goal before checking db (expensive operation) is to perform cheaper operations to handle edge cases
	// Check validity of email address before querying db
	if _, err := mail.ParseAddress(request.Email); err != nil {
		checkError(w, http.StatusBadRequest, "Invalid username or password.")
		log.Print("Invalid email address: ", err)
		return
	}

	// Validate Password complexity before running more expensive bcrypt operations in case of an edge case
	if len(request.Password) < 8 {
		checkError(w, http.StatusBadRequest, "Invalid username or password.")
		log.Print("Password too short. Does not meet complexity requirements of at least 8 characters: ", err)
		return
	}

	if len(request.Password) > 72 {
		checkError(w, http.StatusBadRequest, "Invalid username or password.")
		log.Print("Password too long. Does not meet complexity requirements of less than 72 characters: ", err)
		return
	}

	// Query db to see if this email exists. if it does we want to store the password_hash to compare later
	emailCheck := a.db.QueryRow(r.Context(), "SELECT password_hash FROM users WHERE email = $1", request.Email)
	if err = emailCheck.Scan(&password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Chose generic error messages over specific ones due to possibility of enumeration attacks, and following best security principles
			checkError(w, http.StatusUnauthorized, "Invalid username or password")
			log.Print("Email does not exist in DB: ", err)
			return
		} else {
			checkError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
			log.Print("Error querying users table in Giftwise DB: ", err)
			return
		}

	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			checkError(w, http.StatusUnauthorized, "Invalid username or password.")
			log.Print("Passwords do not match: ", err)
			return
		} else {
			checkError(w, http.StatusUnauthorized, "Invalid username or password.")
			log.Print("Error: ", err)
			return
		}
	}

	// If email exists in DB and password hash from DB matches the password they entered hash, then we write a success response back to user
	writeJSON(w, http.StatusAccepted, "Login successful")

}
