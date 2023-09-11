package handler

import (
	"github.com/tank130701/url-shortener-back-end/internal/services"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
	baseUrl string
}

func NewHandler(services *services.Service, baseUrl string) *Handler {
	return &Handler{
		services: services,
		baseUrl: baseUrl,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("api")
	{
		api.POST("/shorten", h.CreateUrl)
	}
	router.GET("/:shortURL", h.GetFullUrl) 
	
	return router
}
