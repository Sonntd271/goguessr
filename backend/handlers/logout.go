package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	validator "backend/api/validators"
	"backend/shared"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, ss *shared.SessionStorage) {
	// LogoutHandler receives the API token from the header and performs a validation to see if it is an existing API token
	// If it is a valid API token, it removes the existing record, requiring the user to login again
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authBearer := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authBearer, "Bearer ")

	if !validator.ValidateToken(token, ss) {
		http.Error(w, "Token not found", http.StatusBadRequest)
		return
	}

	ss.Lock()
	defer ss.Unlock()

	ss.Sessions[token] = false

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
