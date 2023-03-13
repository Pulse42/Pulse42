package example

import "github.com/gofiber/fiber/v2"

func worldHandler(c *fiber.Ctx) error {
	c.Status(200)
	return c.SendString("World")
}
