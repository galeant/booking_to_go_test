package family

import (
	"encoding/json"
	"latihan/common"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type FamilyHandler struct {
	Service   *FamilyService
	Validator *validator.Validate
}

func (h *FamilyHandler) GetList(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	families, err := h.Service.GetData(idInt)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error", 500)
		return
	}
	common.SuccessResponseMux(w, families, "Success", 200)

}

func (h *FamilyHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	request := CreateFamilyRequest{}
	h.Validator.RegisterValidation("relation", RelationValidation)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}

	if err := h.Validator.Struct(request); err != nil {
		common.ErrorValidationMux(w, err)
		return
	}

	families, err := h.Service.Update(idInt, request)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error", 500)
		return
	}
	common.SuccessResponseMux(w, families, "Success", 200)
}
