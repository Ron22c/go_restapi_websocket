package posts

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func NewPostRouter() http.Handler {
	r := chi.NewRouter();

	r.Get("/", GetAllPosts)

	return r;
}