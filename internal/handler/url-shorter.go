package handler

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tank130701/url-shortener-back-end/internal/helpers"
	"github.com/tank130701/url-shortener-back-end/internal/models"
)

func (h *Handler) CreateUrl(ctx *gin.Context) {
	var url models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	strCatUrl := helpers.RemoveDomainError(url.FullURL)
	shortUrl, err := h.services.CreateUrl(strCatUrl)
	// fmt.Println("Handler")
	// fmt.Println(strCatUrl, shortUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"short_url": "http://www." + h.baseUrl + "/" + shortUrl,
	})

}

func (h *Handler) GetFullUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortURL")
	fullCatUrl, err := h.services.UrlShortener.GetFullUrl(shortUrl)
	fullUrl := "https://www." + fullCatUrl
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Redirect(http.StatusFound, fullUrl)
}
