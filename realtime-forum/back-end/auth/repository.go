package auth

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context) error
}
