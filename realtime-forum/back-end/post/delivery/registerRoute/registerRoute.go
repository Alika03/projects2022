package registerRoute

import (
	"back-end/post"
	"back-end/post/delivery/http/handler"
	"net/http"
)

func RegisterPostHTTPRoute(mux *http.ServeMux, useCase post.UseCase) {
	h := handler.NewHandler(useCase)

	mux.HandleFunc("post/add", h.CreatePost)
	mux.HandleFunc("post/id", h.GetById)
	mux.HandleFunc("post/all", h.GetAll)
}
