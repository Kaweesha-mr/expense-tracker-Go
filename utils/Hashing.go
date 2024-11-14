package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error Hashing the password", err)
		return "", err
	}

	return string(bytes), nil
}

func checkPasswordHash(pasword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pasword))
	return err == nil
}
