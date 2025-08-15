package user

import (
	"encoding/json"
	"fmt"
	"latihan/common"
	"net/http"
	"os/user"

	"strconv"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	paginateParam := r.URL.Query().Get("paginate")
	pageParam := r.URL.Query().Get("page")

	paginate := 10
	page := 1

	if val, err := strconv.Atoi(paginateParam); err == nil && val > 0 {
		paginate = val
	}
	if val, err := strconv.Atoi(pageParam); err == nil && val > 0 {
		page = val
	}

	users, total, err := h.Service.GetDataUser(search, paginate, page)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Data", 500)
		return
	}
	common.SuccessResponseMux(w, map[string]any{
		"users":      users,
		"total_data": total,
	}, 200)
}

func GetListHandler(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	paginateParam := r.URL.Query().Get("paginate")
	pageParam := r.URL.Query().Get("page")

	paginate := 10
	page := 1

	if val, err := strconv.Atoi(paginateParam); err == nil && val > 0 {
		paginate = val
	}
	if val, err := strconv.Atoi(pageParam); err == nil && val > 0 {
		page = val
	}

	users, totalData, err := GetData(search, paginate, page)
	if err != nil {
		common.ErrorResponseMux(w, err, "Error Get Data", 500)
		return
	}

	common.SuccessResponseMux(w, map[string]any{
		"users":      users,
		"total_data": totalData,
	}, 201)

}

func GetDetailHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// fmt.Println("test masuk")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.ErrorResponseMux(w, err, "Invalid request", 422)
		return
	}

	if err := Validate.Struct(user); err != nil {
		fmt.Println(err)
		common.ErrorValidationMux(w, err)
		return
	}

	// Contoh: cetak data ke terminal
	fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)

	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// if err := validate.Struct(req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
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
