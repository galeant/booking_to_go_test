package routes

import (
	"latihan/internal/user"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/user", user.GetListHandler).Methods("GET")
	// r.HandleFunc("/user/{id}", user.GetDetailHandler).Methods("GET")
	r.HandleFunc("/user/create", user.CreateHandler).Methods("POST")
	// r.HandleFunc("/user/update/{id}", user.UpdateHandler).Methods("POST")
	// r.HandleFunc("/user/delete/{id}", user.DeleteHandler).Methods("DELETE")
	// r.GET("/test", auth.TestAja)
}
