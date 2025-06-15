package service

import (
	"GoLang-Notes/jwt-auth-app/model"
	"GoLang-Notes/jwt-auth-app/repository"
	"GoLang-Notes/jwt-auth-app/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string) (string, error) {
	_, exists := repository.GetUserByEmail(email)
	if exists {
		return "", utils.ErrUserExists
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	user := model.User{
		ID:       utils.GenerateUUID(),
		Email:    email,
		Password: string(hashedPassword),
	}

	_ = repository.CreateUser(user)

	token, err := utils.GenerateJWT(user.Email)
	return token, err
}

func LoginUser(email, password string) (string, error) {
	user, exists := repository.GetUserByEmail(email)
	if !exists {
		return "", utils.ErrInvalidCreds
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", utils.ErrInvalidCreds
	}

	return utils.GenerateJWT(user.Email)
}
