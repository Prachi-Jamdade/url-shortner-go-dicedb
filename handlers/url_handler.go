package handlers

import (
	"net/http"
	"url-shortner-dicedb/models"
	"url-shortner-dicedb/services"
	"url-shortner-dicedb/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shorten", CreateShortURL)
	router.GET("/:id", RedirectURL)
}

func CreateShortURL(c *gin.Context) {
	var requestBody models.URL
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	shortURL, err := services.CreateShortURLService(requestBody.LongURL)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error creating short URL")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, gin.H{"short_url": shortURL})
}

func RedirectURL(c *gin.Context) {
	id := c.Param("id")
	longURL, err := services.GetOriginalURLService(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "URL not found")
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
