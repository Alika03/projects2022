package models

import "time"

type Post struct {
	Id        string
	Title     string
	Content   string
	CreatedAt time.Time
}

type PostPagination struct {
	TotalItems int
	PagesCount int
	PerPage    int
	Posts      []*Post
}
