package usecase

import (
	"back-end/auth"
	"back-end/models"
	"back-end/utils"
	"context"
	"github.com/dgrijalva/jwt-go/v4"
)

type UseCase struct {
	repo auth.UserRepository
}

func NewAuthUseCase(repo auth.UserRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) SignUp(ctx context.Context, username, password string) error {
	var model = &models.User{
		Id:       utils.GenerateUuid().String(),
		Username: username,
		Password: password,
	}

	if err := u.repo.CreateUser(ctx, model); err != nil {
		return err
	}

	claimAccess := &jwt.StandardClaims{
		Audience:  model.Id,
		ExpiresAt: nil,
		ID:        "",
		IssuedAt:  nil,
		Issuer:    "",
		NotBefore: nil,
		Subject:   "",
	}
	return nil
}

func (u *UseCase) SignIn(ctx context.Context, username, password string) error {
	return nil
}

func (u *UseCase) ParseToken(ctx context.Context, accessToken string) error {
	return nil
}
