package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	generator "backend/api/generators"
	validator "backend/api/validators"
	"backend/shared"
)

type GuessRequest struct {
	Guess int `json:"guess"`
}

func GuessHandler(w http.ResponseWriter, r *http.Request, ss *shared.SessionStorage, ans *int) {
	// GuessHandler only accepts POST requests containing a guess number
	// It will strip the API token from the request header and validate it
	// If the guess is correct, it will regenerate the hidden number
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

	authBearer := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authBearer, "Bearer ")

	if !validator.ValidateToken(token, ss) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if gr.Guess == *ans {
		hiddenNumber, err := generator.GenerateRandomNumber(10)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		*ans = hiddenNumber
		fmt.Println("New hidden number:", *ans)

		response := map[string]string{"message": "Correct! Generating a new number..."}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]string{"message": "Incorrect guess, please try again"}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
