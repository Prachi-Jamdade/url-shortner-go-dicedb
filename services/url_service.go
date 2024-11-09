package services

import (
	"fmt"
	"url-shortner-dicedb/models"
	"url-shortner-dicedb/repository"

	"github.com/google/uuid"
)

func CreateShortURLService(longURL string) (string, error) {
	id := uuid.New().String()[:8] // Shorten ID for URL

	url := models.URL{
		ID:       id,
		LongURL:  longURL,
		ShortURL: "http://localhost:8080/" + id,
	}

	if err := repository.SaveURL(url); err != nil {
		fmt.Println(err)
		return "", err
	}

	return url.ShortURL, nil
}

func GetOriginalURLService(id string) (string, error) {
	url, err := repository.FindURLByID(id)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}
