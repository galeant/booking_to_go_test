package routes

import (
	"latihan/internal/user"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	userHandler := &user.UserHandler{Service: &user.UserService{}}

	r.HandleFunc("/user", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.GetDetail).Methods("GET")
	r.HandleFunc("/user/create", userHandler.CreateHandler).Methods("POST")
	r.HandleFunc("/user/update/{id}", userHandler.UpdateHandler).Methods("POST")
	r.HandleFunc("/user/delete/{id}", userHandler.DeleteHandler).Methods("DELETE")
	// r.GET("/test", auth.TestAja)
}
