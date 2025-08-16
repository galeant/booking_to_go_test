package nationality

import (
	"net/http"
)

type NationalityHandler struct {
	Service *NationalityService
}

func (h *NationalityHandler) GetList(w http.ResponseWriter, r *http.Request) {
}

func (h *NationalityHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
}

func (h *NationalityHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (h *NationalityHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (h *NationalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
