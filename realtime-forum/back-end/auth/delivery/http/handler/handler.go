package handler

import (
	"back-end/auth"
	"back-end/auth/delivery/http/httpHelper"
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

	// Request
	dto := &AuthDTO{}
	if err := httpHelper.BindJson(request, dto); err != nil {
		_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle request
	if err := h.uc.SignUp(ctx, dto.Username, dto.Password); err != nil {
		_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	httpHelper.JsonCodeResponse(response, http.StatusOK)
}

func (h *Handler) SignIn(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Request
	dto := &AuthDTO{}
	if err := httpHelper.BindJson(request, dto); err != nil {
		_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle request
	tokens, err := h.uc.SignIn(ctx, dto.Username, dto.Password)
	if err != nil {
		_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	response.Header().Add(tokens.TokenType+"-1", tokens.AccessToken)
	response.Header().Add(tokens.TokenType+"-2", tokens.RefreshToken)

	httpHelper.JsonCodeResponse(response, http.StatusOK)
}
