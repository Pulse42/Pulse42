package auth

import (
	"backend/pkg/store"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func logoutHandler(c *fiber.Ctx) error {
	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = session.Destroy()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}
	return c.SendStatus(fiber.StatusOK)
}
