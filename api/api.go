package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handlerGetIndex)
	r.Get("/health", handlerGetHealth)
	r.Get("/{code}", handlerGetShortenUrl)
	r.Post("/api/shorten", handlerPostShortenUrl)

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type PostResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlerGetIndex(w http.ResponseWriter, r *http.Request) {
}

func handlerGetHealth(w http.ResponseWriter, r *http.Request) {
}

func handlerGetShortenUrl(w http.ResponseWriter, r *http.Request) {

}

func handlerPostShortenUrl(w http.ResponseWriter, r *http.Request) {

}
