package main

import (
	"fmt"
	"latihan/config"
	"latihan/internal/user"
	"latihan/routes"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Validate *validator.Validate // variabel global
func main() {
	// Load env
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Gagal memuat file .env")
	}
	// Load DB
	config.ConnectDB()
	// Validation register
	Validate = validator.New()
	Validate.RegisterValidation("relation", user.RelationTypeValidator)
	// Rout
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	// List
	http.ListenAndServe(":8000", r)
}
