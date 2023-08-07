package services

import "github.com/tank130701/url-shortener-back-end/pkg/repository"

type UrlShortener interface{
	CreateUrl(fullURL string)(string, error)
	GetFullUrl(shortUrl string)(string, error)
}


type Service struct{
	UrlShortener
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UrlShortener: NewUrlShortenerService(repo.UrlShortener),
	}
}