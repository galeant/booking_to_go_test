package user

import (
	"encoding/json"
	"net/http"
	"os/user"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetListHandler(w http.ResponseWriter, r *http.Request) {

	users, err := GetData()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Family created successfully",
	})
}

func GetDetailHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req user.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var req user.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
