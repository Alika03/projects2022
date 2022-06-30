package server

import (
	"back-end/auth"
	"back-end/auth/delivery/http/registerRoute"
	"back-end/auth/repository/postgressDb"
	"back-end/auth/usecase"
	"back-end/config"
	"database/sql"
	"fmt"
	"log"
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
	dns := fmt.Sprintf("postgres://postgres:%v@%v:%v/%v?sslmode=disable",
		config.GetConfig().Db.Password,
		config.GetConfig().Db.Host,
		config.GetConfig().Db.Port,
		config.GetConfig().Db.DbName,
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
