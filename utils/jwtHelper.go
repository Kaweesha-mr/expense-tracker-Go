package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string) (string, error) {

	//Define JWT expiration time
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Kaweesha_Mr",
		},
	}

	// Create a new JWT token with the claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Println("Error generating access token:", err)
		return "", err
	}

	return accessToken, nil
}

// GenerateRefreshToken generates a refresh token.
func GenerateRefreshToken(username string) (string, error) {
	// Define JWT expiration time (e.g., 7 days)
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Kaweesha_mr",
		},
	}

	// Create the refresh token with the claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Println("Error generating refresh token:", err)
		return "", err
	}

	return refreshToken, nil
}

// ValidateToken validates the access token.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	// Parse the token and check its validity
	// After
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Println("Error validating token:", err)
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
