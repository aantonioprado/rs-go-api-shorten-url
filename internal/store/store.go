package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	db *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(db *redis.Client) Store {
	return store{db}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = genCode()

		if err := s.db.HGet(ctx, "shorten", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}

			return "", fmt.Errorf("faield to get code from shorten hashmap %w", err)
		}
	}

	if err := s.db.HSet(ctx, "shorten", code, _url).Err(); err != nil {
		return "", fmt.Errorf("faield to set code from shorten hashmap %w", err)
	}

	return code, nil
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	fullURL, err := s.db.HGet(ctx, "shorten", code).Result()

	if err != nil {
		return "", fmt.Errorf("faield to get code from shorten hashmap %w", err)
	}

	return fullURL, nil
}
