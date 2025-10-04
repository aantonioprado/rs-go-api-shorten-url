package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, resp apiResponse, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		sendJSON(w, apiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}
