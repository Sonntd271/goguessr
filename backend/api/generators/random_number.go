package api

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GenerateRandomNumber(numRange int) (int, error) {
	// GenerateRandomNumber generates a random number from 1 to a specified range given as an integer
	// Also returns an error as the second return value
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(numRange)))
	if err != nil {
		return 0, errors.New("Error generating random number")
	}
	return int(randomNumber.Int64()) + 1, nil
}
