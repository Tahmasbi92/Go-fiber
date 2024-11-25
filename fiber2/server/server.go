package server

import (
	"net/http"

	"github.com/Tahmasbi92/fiber2/middleware"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Initialize Fiber
	app := fiber.New()

	// Global Middlewares
	app.Use(middleware.GlobalLogger)

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