package auth

import "context"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) error
	ParseToken(ctx context.Context, accessToken string) error
}
