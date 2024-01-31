package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	generator "backend/api/generators"
	"backend/shared"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ss *shared.SessionStorage) {
	// LoginHandler accepts POST requests containing a username and a password
	// It will perform validation logic to validate the username and password
	// If the validation passes, it will return an API token
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

	if lr.Password == "password" {
		token, err := generator.GenerateToken()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ss.Lock()
		defer ss.Unlock()
		ss.Sessions[token] = true

		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", token))
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
