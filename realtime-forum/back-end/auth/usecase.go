package auth

import (
	"back-end/models"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (*models.Tokens, error)
	VerifyAccessToken(ctx context.Context, accessToken string) (*models.User, error)
}
