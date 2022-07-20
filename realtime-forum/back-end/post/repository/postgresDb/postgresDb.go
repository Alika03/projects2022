package postgresDb

import (
	"back-end/models"
	"context"
	"database/sql"
	"log"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) AddPost(ctx context.Context, model *models.Post) error {
	model.CreatedAt = model.CreatedAt.UTC()

	query := "INSERT INTO post (id, title, content, created_at) VALUES($1, $2, $3, $4);"

	_, err := r.db.ExecContext(ctx, query, model.Id, model.Title, model.Content, model.CreatedAt)

	return err
}

func (r *PostRepository) GetAll(ctx context.Context, limit, offset int) ([]*models.Post, error) {

	var limitSql sql.NullInt32
	if limit > 0 {
		limitSql = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := "SELECT * FROM post ORDER BY created_at DESC LIMIT $1 OFFSET $2"

	rows, err := r.db.QueryContext(ctx, query, limitSql, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var postsModel []*models.Post

	for rows.Next() {
		postModel := &models.Post{}

		if err = rows.Scan(postModel.Id, postModel.Title, postModel.Content, postModel.CreatedAt); err != nil {
			return nil, err
		}

		postsModel = append(postsModel, postModel)
	}

	return postsModel, nil
}

func (r *PostRepository) GetById(ctx context.Context, id string) (*models.Post, error) {
	query := "SELECT * FROM post WHERE id = $1;"

	var model = &models.Post{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(model.Id, model.Title, model.Content, model.CreatedAt)
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

func (r *PostRepository) CountAll(ctx context.Context) (int, error) {
	query := `SELECT count(id) FROM post;`

	var count int

	err := r.db.QueryRowContext(ctx, query).Scan(count)
	switch err {
	case sql.ErrNoRows, nil:
		return count, nil
	default:
		return 0, err
	}
}
