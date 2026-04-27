package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to moniVest API!")
	})

	AuthRoutes(starter)
}