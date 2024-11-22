package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3001"))
}

// Handler
func hello(c *fiber.ctx) error {
	return c.sendstring("Hello, world!")
}
