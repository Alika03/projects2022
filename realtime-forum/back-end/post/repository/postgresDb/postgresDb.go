package postgresDb

import (
	"back-end/models"
	"context"
	"database/sql"
	"log"
	"time"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) AddPost(ctx context.Context, model models.Post) error {
	query := "INSERT INTO post (id, title, content) VALUES($1, $2, $3);"

	_, err := r.db.ExecContext(ctx, query, model.Id, model.Title, model.Content)

	return err
}

func (r *PostRepository) GetAll(ctx context.Context, limit, offset string) ([]*models.Post, error) {
	query := "SELECT * FROM post ORDER BY created_at DESC LIMIT $1 OFFSET $2"

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var postsModel []*models.Post

	for rows.Next() {
		postModel := &models.Post{}
		var createdAt time.Time
		if err = rows.Scan(postModel.Id, postModel.Title, postModel.Content, createdAt); err != nil {
			return nil, err
		}
		postModel.CreatedAt = createdAt.Unix()
		postsModel = append(postsModel, postModel)
	}

	return postsModel, nil
}

func (r *PostRepository) GetById(ctx context.Context, id string) (*models.Post, error) {
	query := "SELECT * FROM post WHERE id = $1;"

	var model = &models.Post{}

	err := r.db.QueryRowContext(ctx, query, id).Scan()
	switch {
	case err == sql.ErrNoRows:
		log.Println("there is no post")
		return model, nil
	case err != nil:
		return nil, err
	default:
		return model, nil
	}
}
