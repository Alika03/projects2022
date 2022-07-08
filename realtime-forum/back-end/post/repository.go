package post

import (
	"back-end/models"
	"context"
)

type PostRepository interface {
	AddPost(ctx context.Context, model *models.Post) error
	GetAll(ctx context.Context, limit, offset string) ([]*models.Post, error)
	GetById(ctx context.Context, id string) (*models.Post, error)
}