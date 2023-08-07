package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB int
}

func NewRedisDB(cfg Config)(*redis.Client, error){
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:	  fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:		  cfg.DB,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("Подключение к Redis успешно:", pong)
	return rdb, nil
}