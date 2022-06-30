package handler

import (
	"back-end/auth"
	"context"
	"net/http"
	"time"
)

type Handler struct {
	uc auth.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) SignUp(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dto := &AuthDTO{}

	if err := h.uc.SignUp(ctx, dto.Username, dto.Password); err != nil {
		return
	}

	// set tokens
	// json response
}

func (h *Handler) SignIn(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dto := &AuthDTO{}

	if err := h.uc.SignIn(ctx, dto.Username, dto.Password); err != nil {
		return
	}

	// set tokens
	// json response
}
