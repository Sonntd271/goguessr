package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var hiddenNumber int
var token string

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GuessRequest struct {
	Token string `json:"token"`
	Guess int    `json:"guess"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Simple authentication, subject to change
	if loginReq.Password == "password" {
		token = generateToken()

		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var guessReq GuessRequest
	err := json.NewDecoder(r.Body).Decode(&guessReq)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if guessReq.Token == token {
		if guessReq.Guess == hiddenNumber {
			hiddenNumber = generateRandomNumber()
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, "Incorrect guess", http.StatusNotAcceptable)
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	hiddenNumber = generateRandomNumber()

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/guess", guessHandler)

	fmt.Println("Server is running on: 8080")
	http.ListenAndServe(":8080", nil)
}
