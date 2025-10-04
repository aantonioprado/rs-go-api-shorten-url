package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullUrl, ok := db[code]

		if !ok {
			sendJSON(w, apiResponse{Error: "code not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, apiResponse{Data: getShortenedURLResponse{FullURL: fullUrl}}, http.StatusOK)
	}
}
