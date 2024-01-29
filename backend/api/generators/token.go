package api

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.New("Error generating token")
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}