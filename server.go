package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	app := fiber.New()

	// Routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello I’m Working",
		})
	})
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{

			"status":  http.StatusOK,
			"heakthy": true,
		})
	})
	app.Listen("3000")
}
