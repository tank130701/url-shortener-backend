package services

import (
	"math/rand"

	"github.com/tank130701/url-shortener-back-end/pkg/repository"
)

type UrlShortenerService struct{
	repo repository.UrlShortener
}

func NewUrlShortenerService(repo repository.UrlShortener) *UrlShortenerService{
	return &UrlShortenerService{repo: repo}
}

func (s *UrlShortenerService) CreateUrl(fullURL string)(string, error){
	shortUrl := generateShortURL()
	err := s.repo.SaveShortUrl(shortUrl, fullURL)
	if err != nil{
		return "", err
	}
	return shortUrl, nil
}

func (s *UrlShortenerService) GetFullUrl(shortUrl string)(string, error){
	return s.repo.GetFullUrl(shortUrl)
}

func generateShortURL() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(charSet))
		result += string(charSet[randomIndex])
	}
	return result
}