package handler

import (
	"back-end/models"
	"back-end/post"
	"context"
	"net/http"
	"time"
)

type Handler struct {
	uc post.UseCase
}

func NewHandler(uc post.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreatePost(response http.Response, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := h.uc.CreatePost(ctx, &models.Post{}); err != nil {
		// handle err json
		return
	}

	// response json
}
