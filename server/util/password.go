package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %v", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	//CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
	//Returns nil on success, or an error on failure.
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
