package auth

import (
	"backend/pkg/store"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func whoAmIHandler(c *fiber.Ctx) error {
	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	name := session.Get("user")
	if name == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not logged in",
		})
	}

	return c.Status(200).SendString("Hello " + name.(string))
}
