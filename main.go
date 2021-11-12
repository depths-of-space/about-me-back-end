package main

import (
	"about-me-back-end/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.HandlerRouterPost(app)

	app.Listen(":3000")
}
