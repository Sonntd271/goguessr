package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	validator "backend/api/validators"
	"backend/shared"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, ss *shared.SessionStorage) {
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
