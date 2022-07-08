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

func (j *JwtRepository) HasAccessTokenById(ctx context.Context, accessId string) (bool, error) {
	var has bool

	query := "SELECT exists(SELECT * FROM access_token WHERE id == $1);"

	row := j.db.QueryRowContext(ctx, query, accessId)
	if err := row.Scan(&has); err != nil {
		return false, err
	}

	return has, nil
}
