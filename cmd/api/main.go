package main

import (
	"aantonioprado/rs-go-api-shorten-url/internal/api"
	"aantonioprado/rs-go-api-shorten-url/internal/store"
	"log/slog"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("Code executed successfully")
}

func run() error {
	// db := make(map[string]string)
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(db)
	handler := api.NewHandler(store)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
