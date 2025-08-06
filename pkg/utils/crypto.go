package utils

import (
	"inorder/pkg/types"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyUser(user *types.User, password string) (bool, error) {
	if user == nil {
		return false, ErrUserNotFound
	}
	return VerifyPassword(user.HashedPassword, password)
}
