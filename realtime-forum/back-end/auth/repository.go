package auth

import (
	"back-end/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, model *models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

type JwtRepository interface {
	AddAccessToken(ctx context.Context, tr *sql.Tx, model *models.AccessToken) error
	AddRefreshToken(ctx context.Context, tr *sql.Tx, model *models.RefreshToken) error
	//GetAccessTokenById(ctx context.Context, accessTokenId string) (*models.Jwt, error)
}
