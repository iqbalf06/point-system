package handlers

import "github.com/gofiber/fiber/v2"

func Handler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"point": "system app",
	})
}
