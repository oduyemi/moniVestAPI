package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(starter *fiber.App) {
	starter.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to moniVest API!")
	})

}