package main

import (
	"GoLang-Notes/jwt-auth-app/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
