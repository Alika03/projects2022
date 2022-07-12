package handler

import (
	"back-end/models"
	"back-end/pkg/httpHelper"
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

func (h *Handler) CreatePost(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := h.uc.CreatePost(ctx, &models.Post{}); err != nil {
		// handle err json
		return
	}

	// response json
}

func (h *Handler) GetAll(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pagination := paginationDTO{}
	paginationModel := &models.Pagination{
		Page:   pagination.Page,
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
	}

	resp, err := h.uc.GetAll(ctx, paginationModel)
	if err != nil {
		_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusConflict)
		return
	}

	// response json
	_ = httpHelper.JsonResponse(response, resp)
}

func (h *Handler) GetById(response http.ResponseWriter, request *http.Request) {

}
