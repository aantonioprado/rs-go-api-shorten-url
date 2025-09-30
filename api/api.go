package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handlerGetIndex(db))
	r.Get("/health", handlerGetHealth(db))
	r.Get("/{code}", handlerGetShortenUrl(db))
	r.Post("/api/shorten", handlerPostShortenUrl(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type PostResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlerGetIndex(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func handlerGetHealth(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func handlerGetShortenUrl(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func handlerPostShortenUrl(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
