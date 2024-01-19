package main

import (
	"point-system/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":8080")
}
