package auth

import (
	"back-end/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}
