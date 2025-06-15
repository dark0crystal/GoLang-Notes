package controller

import (
	"GoLang-Notes/jwt-auth-app/service"
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := service.RegisterUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
