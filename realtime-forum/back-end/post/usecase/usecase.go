package usecase

import (
	"back-end/models"
	"back-end/post"
	"context"
)

type UseCase struct {
	repo post.PostRepository
}

func NewUseCase(repo post.PostRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreatePost(ctx context.Context, model *models.Post) error {
	return nil
}

func (u *UseCase) GetAll(ctx context.Context, pagination models.Pagination) ([]*models.Post, error) {
	return nil, nil
}
func (u *UseCase) GetById(ctx context.Context, id string) (*models.Post, error) {
	return nil, nil
}
