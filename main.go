package main

import (
	"log"
	"url-shortner-dicedb/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	handlers.RegisterRoutes(router)

	log.Fatal(router.Run(":8080"))
}
