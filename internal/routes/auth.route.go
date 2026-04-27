package routes

import (
	"moniVestAPI/internal/controllers"
	"moniVestAPI/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/v1/auth")

	api.Post("/register", controllers.Register)
	api.Post("/verify-otp", controllers.VerifyOTP)
	api.Post("/login", controllers.Login)
	api.Post("/refresh-token", controllers.RefreshToken)
	api.Post("/resend-otp", controllers.ResendOTP)
	api.Post("/logout", middleware.Protected(), controllers.Logout)
}