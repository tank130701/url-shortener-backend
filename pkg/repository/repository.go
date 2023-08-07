package repository

import "github.com/redis/go-redis/v9"

type UrlShortener interface{
	SaveShortUrl(shortURL, fullURL string) error
	GetFullUrl(shortURL string) (string, error)
}

type Repository struct {
	UrlShortener
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{
		UrlShortener: NewSaveUrl(rdb),
	}
}