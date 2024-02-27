package AuthUtils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePhoneNumber(phone_number string) error {
	if len(phone_number) != 11 {
		errorMessage := "invalid phone number"
		fmt.Println("Error:", errorMessage)
		return errors.New(errorMessage)
	}

	// should also check here to see if the phone number exists already
	return nil
}

func generatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidateAndHashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password must not be empty")
	}

	if len(password) <= 5 {
		return "", errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		return "", errors.New("error generating password hash")
	}

	return hashedPassword, nil
}

func GenerateToken(phone_number string, hashedPassword string) (string, error) {
	// Concatenate phone_number and hashedPassword to create a basic token
	basicToken := fmt.Sprintf("%s:%s", phone_number, hashedPassword)

	// Hash the basic token using SHA-256
	hashedToken := sha256.Sum256([]byte(basicToken))

	// Convert the hashed token to a hexadecimal string
	tokenString := hex.EncodeToString(hashedToken[:])

	return tokenString, nil
}
