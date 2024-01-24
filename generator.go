package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"math/big"
)

func generateRandomNumber() int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatal("Error generating random number")
	}
	return int(randomNumber.Int64()) + 1
}

func generateToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Error generating token")
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
