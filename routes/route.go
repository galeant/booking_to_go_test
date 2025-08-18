package routes

import (
	"latihan/internal/family"
	"latihan/internal/nationality"
	"latihan/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {

	validator := validator.New()
	validator.RegisterValidation("relation", family.RelationValidation)

	userHandler := &user.UserHandler{
		Service:   &user.UserService{DB: db},
		Validator: validator,
	}

	familyHandler := &family.FamilyHandler{
		Service:   &family.FamilyService{DB: db},
		Validator: validator,
	}

	nationalityHandler := &nationality.NationalityHandler{
		Service:   &nationality.NationalityService{DB: db},
		Validator: validator,
	}

	// User
	r.HandleFunc("/user", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.GetDetail).Methods("GET")
	r.HandleFunc("/user/create", userHandler.CreateHandler).Methods("POST")
	r.HandleFunc("/user/{id}/update", userHandler.UpdateHandler).Methods("POST")
	r.HandleFunc("/user/{id}/delete", userHandler.DeleteHandler).Methods("DELETE")

	// Family
	r.HandleFunc("/user/{id}/family", familyHandler.GetList).Methods("GET")
	r.HandleFunc("/user/{id}/family/update", familyHandler.Update).Methods("POST")

	// Nationality
	r.HandleFunc("/nationality", nationalityHandler.GetList).Methods("GET")
	r.HandleFunc("/nationality/{id}", nationalityHandler.GetDetail).Methods("GET")
	r.HandleFunc("/nationality/create", nationalityHandler.Create).Methods("POST")
	r.HandleFunc("/nationality/update/{id}", nationalityHandler.Update).Methods("POST")
	r.HandleFunc("/nationality/delete/{id}", nationalityHandler.Delete).Methods("DELETE")
}
