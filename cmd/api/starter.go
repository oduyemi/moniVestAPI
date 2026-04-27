package main

import (
	"log"
	"moniVestAPI/internal/config"
	"moniVestAPI/internal/models"
	"moniVestAPI/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	config.DbConnect()

	err := models.CreateUserIndexes(config.UserCollection)
	if err != nil {
		log.Fatal("Failed to create indexes:", err)
	}

	starter := fiber.New()
	routes.SetupRoutes(starter)
	starter.Listen(":3000")
}