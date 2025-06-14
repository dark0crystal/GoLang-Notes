package main

import (
	"GoLang-Notes/clean-architecture/internal/repository"
	"GoLang-Notes/clean-architecture/internal/usecase"
	"fmt"

	"net/http"
)

func main() {
	repo := repository.NewInMemoryUserRepo()
	uc := usecase.NewUserUseCase(repo)
	handler := http.NewUserHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
