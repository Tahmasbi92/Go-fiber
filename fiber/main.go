package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initiolize Fiber
	app := fiber.New()

	// Routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello Everything is working perfect and as expected!",
			"status":  http.StatusOK,
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
