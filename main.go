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

	var lr LoginRequest
	err := json.NewDecoder(r.Body).Decode(&lr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Simple authentication, subject to change
	if lr.Password == "password" {
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

	var gr GuessRequest
	err := json.NewDecoder(r.Body).Decode(&gr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if gr.Token == token {
		if gr.Guess == hiddenNumber {
			hiddenNumber = generateRandomNumber()
			w.Header().Add("Content-Type", "text/plain")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Correct! Generating a new number..."))
		} else {
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte("Incorrect guess, please try again"))
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	hiddenNumber = generateRandomNumber()

	http.Handle("/", CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login":
			loginHandler(w, r)
		case "/guess":
			guessHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})))

	fmt.Println("Server is running on: 8080")
	http.ListenAndServe(":8080", nil)
}
