package api

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"

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

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	const n = 8
	bytes := make([]byte, n)
	for i := range n {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)

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
		code := chi.URLParam(r, "code")
		data, ok := db[code]

		if !ok {
			sendJSON(
				w,
				Response{Error: "URL not found"},
				http.StatusNotFound,
			)
			return
		}

		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}

func handlerPostShortenUrl(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(
				w,
				Response{Error: "invalid request body"},
				http.StatusBadRequest,
			)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid URL passed"}, http.StatusBadRequest)
			return
		}

		code := genCode()
		db[code] = body.URL

		sendJSON(
			w,
			Response{Data: map[string]string{"code": code}},
			http.StatusOK,
		)
	}
}
