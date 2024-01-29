package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	generator "backend/api/generators"
	validator "backend/api/validators"
	"backend/types"
)

type GuessRequest struct {
	Token string `json:"token"`
	Guess int    `json:"guess"`
}

func GuessHandler(w http.ResponseWriter, r *http.Request, ss *types.SessionStorage, ans *int) {
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

	if !validator.ValidateToken(gr.Token, ss) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if gr.Guess == *ans {
		hiddenNumber, err := generator.GenerateRandomNumber(10)
		if err != nil {
			return
		}
		*ans = hiddenNumber
		fmt.Println("New hidden number:", *ans)

		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Correct! Generating a new number..."))
	} else {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("Incorrect guess, please try again"))
	}
}
