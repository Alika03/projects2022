package handler

type postDTO struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type paginationDTO struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
