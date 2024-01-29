package handlers

import (
	"encoding/json"
	"net/http"

	generator "backend/api/generators"
	"backend/types"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ss *types.SessionStorage) {
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
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
