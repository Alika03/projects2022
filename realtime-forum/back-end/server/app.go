package server

import (
	"back-end/auth"
	registerAuthRoute "back-end/auth/delivery/http/registerRoute"
	"back-end/auth/repository/postgressDb"
	authUC "back-end/auth/usecase"
	"back-end/config"
	"back-end/post"
	registerPostRoute "back-end/post/delivery/registerRoute"
	"back-end/post/repository/postgresDb"
	postUC "back-end/post/usecase"
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
	postUC post.UseCase
}

func NewApp() *App {
	db := initDb()
	privateKey := config.GetKeys().Private

	return &App{
		httpServer: nil,
		authUC:     authUC.NewAuthUseCase(postgressDb.NewUserRepository(db), postgressDb.NewJwtRepository(db), privateKey),
		postUC:     postUC.NewPostUseCase(postgresDb.NewPostRepository(db)),
	}
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	registerAuthRoute.RegisterAuthHTTPRoute(mux, a.authUC)
	registerPostRoute.RegisterPostHTTPRoute(mux, a.postUC)

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
