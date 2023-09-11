package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/tank130701/url-shortener-back-end/internal/app"
	"github.com/tank130701/url-shortener-back-end/internal/handler"
	"github.com/tank130701/url-shortener-back-end/internal/repository"
	"github.com/tank130701/url-shortener-back-end/internal/repository/redis_storage"
	"github.com/tank130701/url-shortener-back-end/internal/repository/sqlite_storage"
	"github.com/tank130701/url-shortener-back-end/internal/services"
)

func main(){

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	
	repo, err := NewStorage()
	if err != nil {
		logrus.Fatalf("failed to initialize storage: %s", err.Error())
	}


	services := services.NewService(repo)
	handlers := handler.NewHandler(services, os.Getenv("DOMAIN") )

	srv := new(app.App)
	go func() {
		if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("UrlShortener Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("UrlShortener Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	// if err := db.Close(); err != nil {
	// 	logrus.Errorf("error occured on db connection close: %s", err.Error())
	// }


}

func NewStorage() (*repository.Repository, error) {
	storageType := os.Getenv("STORAGE")
	switch storageType {
	case "sqlite":
		db, err := sqlite_storage.NewSqliteDB("storage/storage.db")
		if err != nil {
			return nil, err
		}
		return repository.NewRepositorySqlite(db), nil
	case "redis":
		db, err := redis_storage.NewRedisDB(redis_storage.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       0,
		})
		if err != nil {
			return nil, err
		}
		return repository.NewRepositoryRedis(db), nil
	default:
		return nil, errors.New("unsupported storage type")
	}
}
	