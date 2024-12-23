package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func hashRefreshToken(refresh string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash refresh token")
	}
	return string(bytes), nil
}
