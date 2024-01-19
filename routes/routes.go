package routes

import (
	"point-system/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handlers.Handler)
}
