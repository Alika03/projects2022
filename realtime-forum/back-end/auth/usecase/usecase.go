package usecase

import (
	"back-end/auth"
	"back-end/models"
	"context"
)

type UseCase struct {
	repo auth.UserRepository
}

func NewAuthUseCase(repo auth.UserRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) SignUp(ctx context.Context, username, password string) error {
	var model = &models.User{
		Id:       "",
		Username: username,
		Password: password,
	}

	if err := u.repo.CreateUser(ctx, model); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) SignIn(ctx context.Context, username, password string) error {
	return nil
}

func (u *UseCase) ParseToken(ctx context.Context, accessToken string) error {
	return nil
}
