package post

import (
	"back-end/models"
	"context"
)

type UseCase interface {
	CreatePost(ctx context.Context, model *models.Post) error
	GetAll(ctx context.Context) ([]*models.Post, error)
	GetById(ctx context.Context, id string) (*models.Post, error)
}
