package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tank130701/url-shortener-back-end/internal/helpers"
	"github.com/tank130701/url-shortener-back-end/internal/models"
)

func (h *Handler) CreateUrl(c *gin.Context) {
	var url models.URL
		if err := c.ShouldBindJSON(&url); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid data format.",
			})
			return 
		}
		strCatUrl := helpers.RemoveDomainError(url.FullURL)
		shortUrl, err := h.services.CreateUrl(strCatUrl)
		fmt.Println("Handler")
		fmt.Println(strCatUrl, shortUrl)
		if err != nil{
			c.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"short_url":  "http://www." + h.baseUrl + "/" + shortUrl,
		})
		
}

func (h *Handler) GetFullUrl(c *gin.Context) {
	shortUrl := c.Param("shortURL")
	fullCatUrl, err := h.services.UrlShortener.GetFullUrl(shortUrl)
	fullUrl := "https://www." + fullCatUrl
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Short URL not found",
		})
		return
	}
	c.Redirect(http.StatusFound, fullUrl)
}

