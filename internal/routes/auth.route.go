package routes

import (
	"moniVestAPI/internal/handlers"
	"moniVestAPI/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/v1/auth")

	api.Post("/register", handlers.Register)
	api.Post("/verify-otp", handlers.VerifyOTP)
	api.Post("/login", handlers.Login)
	api.Post("/refresh-token", handlers.RefreshToken)
	api.Post("/resend-otp", handlers.ResendOTP)
	api.Post("/logout", middlewares.Protected(), controllers.Logout)
}