package main

import (
	"latihan/config"
	"latihan/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()
	// config.DB.AutoMigrate(&user.User{})
	// seedDummyUser()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)
}
