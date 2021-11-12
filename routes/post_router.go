package routes

import (
	"github.com/gofiber/fiber/v2"
)

//HandlerRouter handler routes POST
func HandlerRouterPost(app *fiber.App) {

	app.Get("/posts", func(c *fiber.Ctx) error {

		message := "Test"

		return c.SendString(message)
	})
}
