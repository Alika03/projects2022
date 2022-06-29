package postgressDb

import (
	"back-end/models"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, username, password string) error {
	return nil
}

func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}
