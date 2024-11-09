package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"url-shortner-dicedb/models"

	"github.com/dicedb/dicedb-go" // DiceDB Go SDK
)

var db *dicedb.Client

func init() {
	db = dicedb.NewClient(&dicedb.Options{
		Addr: "localhost:7379", 
	})
}

// SaveURL stores a URL in DiceDB
func SaveURL(url models.URL) error {
	ctx := context.Background()

	// Serialize the URL struct to JSON
	urlData, err := json.Marshal(url)
	if err != nil {
		return fmt.Errorf("failed to serialize URL: %w", err)
	}

	// Store the serialized URL JSON in DiceDB
	statusCmd := db.Set(ctx, url.ID, urlData, 0)
	_, err = statusCmd.Result()
	if err != nil {
		return fmt.Errorf("failed to save URL: %w", err)
	}
	return nil
}

// FindURLByID retrieves a URL by its ID
func FindURLByID(id string) (models.URL, error) {
	ctx := context.Background()

	// Retrieve the serialized URL data from DiceDB
	stringCmd := db.Get(ctx, id)
	urlData, err := stringCmd.Result()
	if err != nil {
		if err == dicedb.Nil {
			return models.URL{}, errors.New("URL not found")
		}
		return models.URL{}, fmt.Errorf("failed to retrieve URL: %w", err)
	}

	// Deserialize the JSON data back into the URL struct
	var url models.URL
	err = json.Unmarshal([]byte(urlData), &url)
	if err != nil {
		return models.URL{}, fmt.Errorf("failed to decode URL data: %w", err)
	}
	return url, nil
}
