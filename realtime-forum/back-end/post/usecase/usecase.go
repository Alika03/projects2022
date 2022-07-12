package usecase

import (
	"back-end/models"
	"back-end/post"
	"back-end/utils"
	"context"
	"time"
)

type UseCase struct {
	repo post.PostRepository
}

func NewPostUseCase(repo post.PostRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreatePost(ctx context.Context, model *models.Post) error {
	model.Id = utils.GenerateUuid().String()
	model.CreatedAt = time.Now()

	return u.repo.AddPost(ctx, model)
}

func (u *UseCase) GetAll(ctx context.Context, pagination models.Pagination) ([]*models.Post, error) {

	return u.repo.GetAll(ctx, pagination.Limit, pagination.Offset)
}
func (u *UseCase) GetById(ctx context.Context, id string) (*models.Post, error) {
	return u.repo.GetById(ctx, id)
}
