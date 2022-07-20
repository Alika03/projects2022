package usecase

import (
	"back-end/models"
	"back-end/pkg"
	"back-end/post"
	"context"
	"time"
)

const (
	limit = 10
)

type UseCase struct {
	repo post.PostRepository
}

func NewPostUseCase(repo post.PostRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreatePost(ctx context.Context, model *models.Post) error {
	model.Id = pkg.GenerateUuid().String()
	model.CreatedAt = time.Now()

	return u.repo.AddPost(ctx, model)
}

func (u *UseCase) GetAll(ctx context.Context, pagination models.Pagination) (*models.PostPagination, error) {
	totalItems, err := u.repo.CountAll(ctx)
	if err != nil {
		return nil, err
	}

	if totalItems == 0 {
		return &models.PostPagination{}, nil
	}

	if pagination.Limit == 0 {
		pagination.Limit = limit
	}

	computedPagination := pkg.NewPagination(totalItems, pagination.Limit)

	posts, err := u.repo.GetAll(ctx, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, err
	}

	return &models.PostPagination{
		TotalItems: totalItems,
		PagesCount: computedPagination.PagesCount(),
		PerPage:    computedPagination.PerPage(),
		Posts:      posts,
	}, nil
}
func (u *UseCase) GetById(ctx context.Context, id string) (*models.Post, error) {
	return u.repo.GetById(ctx, id)
}
