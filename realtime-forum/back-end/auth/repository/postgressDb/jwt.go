package postgressDb

import (
	"back-end/models"
	"context"
	"database/sql"
)

type JwtRepository struct {
	db *sql.DB
}

func NewJwtRepository(db *sql.DB) *JwtRepository {
	return &JwtRepository{db: db}
}

func (j *JwtRepository) AddAccessToken(ctx context.Context, _ *sql.Tx, model *models.AccessToken) error {

	query := "INSERT INTO access_token (id, user_id, expired_at) VALUES($1, $2, $3);"

	_, err := j.db.ExecContext(ctx, query, model.Id, model.UserId, model.ExpiredAt.UTC())
	if err != nil {
		return err
	}

	return nil
}

func (j *JwtRepository) AddRefreshToken(ctx context.Context, _ *sql.Tx, model *models.RefreshToken) error {
	query := "INSERT INTO refresh_token (id, access_token_id, expired_at) VALUES($1, $2, $3);"

	_, err := j.db.ExecContext(ctx, query, model.Id, model.AccessTokenId, model.ExpiredAt.UTC())
	if err != nil {
		return err
	}

	return nil
}
