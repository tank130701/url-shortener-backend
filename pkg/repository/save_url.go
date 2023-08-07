package repository

import (
	"time"
	"context"
	"github.com/redis/go-redis/v9"
)

type SaveUrl struct{
	rdb *redis.Client
}

func NewSaveUrl(rdb *redis.Client) *SaveUrl {
	return &SaveUrl{rdb: rdb}
}

func(r *SaveUrl) SaveShortUrl(shortURL, fullURL string) error {
	ctx := context.Background()
	// Установка ключа "короткая ссылка" и значения "полная ссылка" в Redis.
	// Устанавливаем срок хранения в один час (например).
	err := r.rdb.Set(ctx, shortURL, fullURL, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func(r *SaveUrl) GetFullUrl(shortURL string) (string, error) {
	ctx := context.Background()
	value, err := r.rdb.Get(ctx, shortURL).Result()
	if err != nil {
		return "", err 
	}
	return value, nil
}