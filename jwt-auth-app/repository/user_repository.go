package repository

import (
	"GoLang-Notes/jwt-auth-app/model"
)

var users = map[string]model.User{}

func CreateUser(user model.User) error {
	users[user.Email] = user
	return nil
}

func GetUserByEmail(email string) (model.User, bool) {
	user, exists := users[email]
	return user, exists
}
