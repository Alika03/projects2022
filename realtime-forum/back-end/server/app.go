package server

import (
	"back-end/auth"
	"back-end/auth/delivery/http/registerRoute"
	"back-end/auth/repository/postgressDb"
	"back-end/auth/usecase"
	"back-end/config"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
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
	privateKey := config.GetKeys().Private

	return &App{
		httpServer: nil,
		authUC: usecase.NewAuthUseCase(
			postgressDb.NewUserRepository(db),
			postgressDb.NewJwtRepository(db),
			privateKey,
		),
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

	log.Println("server launched: ", a.httpServer.Addr)
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func initDb() *sql.DB {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetConfig().Db.User,
		config.GetConfig().Db.Password,
		config.GetConfig().Db.Host,
		config.GetConfig().Db.Port,
		config.GetConfig().Db.DbName,
	)

	db, err := sql.Open("pgx", dns)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
