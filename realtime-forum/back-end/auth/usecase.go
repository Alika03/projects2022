package auth

import (
	"back-end/models"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (*models.Tokens, error)
	ParseToken(ctx context.Context, accessToken string) error
}
