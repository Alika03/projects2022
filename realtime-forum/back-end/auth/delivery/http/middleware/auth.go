package middleware

import (
	"back-end/auth"
	"back-end/pkg/httpHelper"
	"context"
	"net/http"
	"time"
)

const userIdHeaderKey = "user_id"

type AuthMiddleware struct {
	uc auth.UseCase
}

func NewAuthMiddleware(uc auth.UseCase) *AuthMiddleware {
	return &AuthMiddleware{uc: uc}
}

func (a *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		accessToken := request.Header.Get("bearer-1")
		if accessToken == "" {
			_ = httpHelper.ErrJsonResponse(response, auth.ErrUnauthorized.Error(), http.StatusUnauthorized)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userModel, err := a.uc.VerifyAccessToken(ctx, accessToken)
		if err != nil {
			_ = httpHelper.ErrJsonResponse(response, err.Error(), http.StatusUnauthorized)
			return
		}

		request.Header.Add(userIdHeaderKey, userModel.Id)

		next.ServeHTTP(response, request)
	}
}
