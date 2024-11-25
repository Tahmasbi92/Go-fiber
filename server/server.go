package server

import (
	"net/http"

	"github.com/Tahmasbi92/middleware-tutorial/middleware"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Initialize Fiber
	app := fiber.New()

	// Global Middlewares
	app.Use(middleware.RequestLogger)
	// Define Routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello I'm Working",
		})
	})
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"healthy": true,
		})
	})
	app.Listen(":3000")
}
