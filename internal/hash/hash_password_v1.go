package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordV1(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHashV1(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("Error comparing passwords: %s\n", err)
		return false
	}
	return true
}
