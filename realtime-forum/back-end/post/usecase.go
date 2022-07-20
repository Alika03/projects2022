package post

import (
	"back-end/models"
	"context"
)

type UseCase interface {
	CreatePost(ctx context.Context, model *models.Post) error
	GetAll(ctx context.Context, pagination models.Pagination) (*models.PostPagination, error)
	GetById(ctx context.Context, id string) (*models.Post, error)
}
