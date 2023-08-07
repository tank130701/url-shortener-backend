package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tank130701/url-shortener-back-end/pkg/models"
)

const baseUrl = "http://localhost:3000/"

func (h *Handler) CreateUrl(c *gin.Context) {
	var url models.URL
		if err := c.ShouldBindJSON(&url); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid data format.",
			})
			return 
		}
		shortUrl, err := h.services.CreateUrl(url.FullURL)
		if err != nil{
			c.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"short_url":  baseUrl + shortUrl,
		})
}

func (h *Handler) GetFullUrl(c *gin.Context) {
	shortUrl := c.Param("shortURL")
	fullUrl, err := h.services.UrlShortener.GetFullUrl(shortUrl)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Short URL not found",
		})
		return
	}
	c.Redirect(http.StatusFound, fullUrl)
}

