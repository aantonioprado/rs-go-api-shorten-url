package api

import (
	"aantonioprado/rs-go-api-shorten-url/internal/store"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
)

type shorenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string
}

func handleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shorenURLRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, apiResponse{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code, err := store.SaveShortenedURL(r.Context(), body.URL)

		if err != nil {
			slog.Error("failed to create code", "error", err)
			sendJSON(w, apiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
