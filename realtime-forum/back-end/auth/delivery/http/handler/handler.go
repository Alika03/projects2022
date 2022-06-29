package handler

import (
	"back-end/auth"
	"net/http"
)

type Handler struct {
	uc auth.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) SignUp(response http.ResponseWriter, request *http.Request) {
}

func (h *Handler) SignIn(response http.ResponseWriter, request *http.Request) {

}
