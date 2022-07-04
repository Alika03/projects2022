package registerRoute

import (
	"back-end/auth"
	"back-end/auth/delivery/http/handler"
	"net/http"
)

func RegisterAuthHTTPRoute(mux *http.ServeMux, useCase auth.UseCase) {
	h := handler.NewHandler(useCase)

	mux.HandleFunc("/auth/sign-up", h.SignUp)
	mux.HandleFunc("/auth/sign-in", h.SignIn)
}
