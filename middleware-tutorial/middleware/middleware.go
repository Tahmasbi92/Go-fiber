package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(ctx *fiber.Ctx) error {
	log.Println("The User has made a request...")
	return ctx.Next()
}

func TimeLogger(ctx *fiber.Ctx) error {
	hour, minute, second := time.Now().Clock()

	log.Printf("The time is %d:%d:%d", hour, minute, second)

	return ctx.Next()
}
