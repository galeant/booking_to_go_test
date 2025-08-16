package nationality

import (
	"encoding/json"
	"fmt"
	"latihan/common"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

type NationalityHandler struct {
	Service *NationalityService
}

func (h *NationalityHandler) GetList(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	fmt.Println("qdw")
	nationalities, err := h.Service.GetList(search)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Data", 500)
		return
	}
	common.SuccessResponseMux(w, nationalities, "", 200)
}

func (h *NationalityHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	nationality, err := h.Service.GetDetail(idInt)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Data", 500)
		return
	}
	common.SuccessResponseMux(w, nationality, "", 200)
}

func (h *NationalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var nationality Nationality
	if err := json.NewDecoder(r.Body).Decode(&nationality); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}

	if err := validate.Struct(nationality); err != nil {
		common.ErrorValidationMux(w, err)
		return
	}

	nationality, err := h.Service.Create(nationality)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Create Data", 500)
		return
	}
	common.SuccessResponseMux(w, nationality, "Success Create", 200)
}

func (h *NationalityHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // Ambil semua parameter URL
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var nationality Nationality
	if err := json.NewDecoder(r.Body).Decode(&nationality); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}

	if err := validate.Struct(nationality); err != nil {
		common.ErrorValidationMux(w, err)
		return
	}
	fmt.Println(nationality)
	nationality, errDB := h.Service.Update(idInt, nationality)
	if errDB != nil {
		common.ErrorResponseMux(w, errDB, "Error Get Detail", 500)
		return
	}
	common.SuccessResponseMux(w, nationality, "Update Success", 200)
}

func (h *NationalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // Ambil semua parameter URL
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	nationality, err := h.Service.Delete(idInt)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Detail", 500)
		return
	}
	common.SuccessResponseMux(w, nationality, "Delete Success", 200)
}
