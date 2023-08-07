package handler

import (
	"github.com/tank130701/url-shortener-back-end/pkg/services"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
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
