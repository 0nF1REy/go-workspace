package app

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]User{}

func Register(username, password string) error {

	if _, exists := users[username]; exists {
		return errors.New("usuário já existe")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Password: string(hash),
	}

	users[username] = user

	return nil
}

func Authenticate(username, password string) bool {

	user, exists := users[username]
	if !exists {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}

func GetEncryptedPassword(username string) (string, bool) {
	user, exists := users[username]
	if !exists {
		return "", false
	}

	return user.Password, true
}
