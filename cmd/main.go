package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/joho/godotenv"
	urlshortenerbackend "github.com/tank130701/url-shortener-back-end"
	"github.com/tank130701/url-shortener-back-end/pkg/handler"
	"github.com/tank130701/url-shortener-back-end/pkg/repository"
	"github.com/tank130701/url-shortener-back-end/pkg/services"
	"github.com/sirupsen/logrus"
	"context"
)

func main(){

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	
	db, err := repository.NewRedisDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:   0,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := services.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(urlshortenerbackend.Server)
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

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}


}