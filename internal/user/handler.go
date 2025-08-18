package user

import (
	"encoding/json"
	"latihan/common"
	"net/http"

	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service   *UserService
	Validator *validator.Validate
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	// paginateParam := r.URL.Query().Get("paginate")
	// pageParam := r.URL.Query().Get("page")

	// paginate := 10
	// page := 1

	// if val, err := strconv.Atoi(paginateParam); err == nil && val > 0 {
	// 	paginate = val
	// }
	// if val, err := strconv.Atoi(pageParam); err == nil && val > 0 {
	// 	page = val
	// }

	users, err := h.Service.GetData(search)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Data", 500)
		return
	}
	common.SuccessResponseMux(w, users, "", 200)
}

func (h *UserHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetDetail(idInt)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Detail", 500)
		return
	}
	common.SuccessResponseMux(w, user, "", 200)
}

func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	request := UserCreateRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}

	if err := h.Validator.Struct(request); err != nil {
		common.ErrorValidationMux(w, err)
		return
	}
	user, err := h.Service.Create(request)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Create Data", 500)
		return
	}
	common.SuccessResponseMux(w, user, "Success Create", 200)
}

func (h *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	request := UserCreateRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}
	if err := h.Validator.Struct(request); err != nil {
		common.ErrorValidationMux(w, err)
		return
	}
	user, errDB := h.Service.Update(idInt, request)
	if errDB != nil {
		common.ErrorResponseMux(w, errDB, "Error Get Detail", 500)
		return
	}
	common.SuccessResponseMux(w, user, "Success Update", 200)
}

func (h *UserHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Delete(idInt)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Detail", 500)
		return
	}
	common.SuccessResponseMux(w, user, "Delete Success", 200)
}
