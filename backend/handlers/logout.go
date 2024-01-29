package handlers

import (
	"encoding/json"
	"net/http"

	"backend/types"
)

type LogoutRequest struct {
	Token string `json:"token"`
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, ss *types.SessionStorage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var lr LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&lr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	ss.Lock()
	defer ss.Unlock()

	ss.Sessions[lr.Token] = false

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
