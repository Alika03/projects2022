package server

import (
	"back-end/auth"
	"back-end/auth/delivery/http/registerRoute"
	"back-end/auth/repository/postgressDb"
	"back-end/auth/usecase"
	"database/sql"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
}

func NewApp() *App {
	db := initDb()

	return &App{
		httpServer: nil,
		authUC: usecase.NewAuthUseCase(
			postgressDb.NewUserRepository(db)),
	}
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	registerRoute.RegisterAuthHTTPRoute(mux, a.authUC)

	a.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func initDb() *sql.DB {
	return nil
}
