package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Define the secret key
	secretKey := []byte("your-256-bit-secret")

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	// Print the token
	fmt.Println("Generated Token:", tokenString)
}
