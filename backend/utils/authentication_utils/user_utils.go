// userUtils.go
package userUtils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePhoneNumber(phone_number string) error {
	if len(phone_number) != 11 {
		return errors.New("invalid phone number")
	}
	return nil
}

func generatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidateAndHashPassword(password string) (string, bool, error) {
	if password == "" {
		return "", false, errors.New("password must not be empty")
	}

	if len(password) <= 5 {
		return "", false, errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		return "", false, errors.New("error generating password hash")
	}

	return hashedPassword, true, nil
}
