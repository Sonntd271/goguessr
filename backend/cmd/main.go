package main

import (
	"fmt"
	"log"
	"net/http"

	generator "backend/api/generators"
	"backend/handlers"
	"backend/shared"
)

func main() {
	sessionStorage := shared.SessionStorage{
		Sessions: make(map[string]bool),
	}

	hiddenNumber, err := generator.GenerateRandomNumber(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hidden number:", hiddenNumber)

	http.Handle("/", handlers.CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login":
			handlers.LoginHandler(w, r, &sessionStorage)
		case "/logout":
			handlers.LogoutHandler(w, r, &sessionStorage)
		case "/guess":
			handlers.GuessHandler(w, r, &sessionStorage, &hiddenNumber)
		default:
			http.NotFound(w, r)
		}
	})))

	fmt.Println("Server is running on: 8080")
	http.ListenAndServe(":8080", nil)
}
