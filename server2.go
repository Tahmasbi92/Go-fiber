package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func run() {
	// Initialize Fiber
	app := fiber.New()

	// Global Middlewares
	app.Use(middleware.GlobalLogger)

	// Define Routes
	app.Get("/request-logger", middleware.RequestLogger, func(ctx *fiber.ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Request log is printed out",
		})
	})
	app.Get("/tima-logger", middleware.TimeLogger, func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Time log is printed out",
		})
	})

	app.Listen(":3000")
}
