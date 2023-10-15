package redis_storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct{
	rdb *redis.Client
}

func NewRedisStorage(rdb *redis.Client) *RedisStorage {
	return &RedisStorage{rdb: rdb}
}

var ctx = context.Background()

func(r *RedisStorage) SaveShortUrl(shortURL, fullURL string) error {
	
	// Установка ключа "короткая ссылка" и значения "полная ссылка" в Redis.
	// Устанавливаем срок хранения в один час (например).
	// checkUrl := r.rdb.
	err := r.rdb.Set(ctx, shortURL, fullURL, time.Hour).Err()
	if err != nil {
		return fmt.Errorf("error saving url in database: %w", err)
	}
	return nil
}



func(r *RedisStorage) GetFullUrl(shortURL string) (string, error) {
	value, err := r.rdb.Get(ctx, shortURL).Result()
	if err != nil {
		return "", fmt.Errorf("error getting url from database: %w", err)
	}
	return value, nil
}