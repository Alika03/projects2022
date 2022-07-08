package postgressDb

import (
	"back-end/models"
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, model *models.User) error {
	query := "insert into users_user (id, username, hash_password) values($1, $2, $3)"

	_, err := u.db.ExecContext(ctx, query, model.Id, model.Username, model.HashPassword)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := "select * from users_user where username = $1"

	var model = &models.User{}

	row := u.db.QueryRowContext(ctx, query, username)

	err := row.Scan(&model.Id, &model.Username, &model.HashPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no such username")
	case err != nil:
		return nil, err
	default:
		return model, nil
	}
}

func (u *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	query := "select * from users_user where id = $1"

	var model = &models.User{}

	row := u.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&model.Id, &model.Username, &model.HashPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no such username")
	case err != nil:
		return nil, err
	default:
		return model, nil
	}
}
