package auth

import (
	"backend/pkg/interfaces"
	"backend/pkg/models"
	"backend/pkg/store"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func loginHandler(c *fiber.Ctx) error {
	credentials := new(interfaces.AuthInterface)

	if err := c.BodyParser(credentials); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing data",
			"error":   err.Error(),
		})
	}

	user := new(models.UserModel)
	_, err := user.GetUserByMail(credentials.Mail)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusNotFound, err)
	}

	err = user.ComparePassword(credentials.Password)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusUnauthorized, err)
	}

	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	session.Set("user", user.Username)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = session.Save()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
