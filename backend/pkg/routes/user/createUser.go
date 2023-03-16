package user

import "C"
import (
	"backend/pkg/models"
	"backend/pkg/utils"
	"backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func createUserHandler(c *fiber.Ctx) error {
	user := new(models.UserModel)

	if err := c.BodyParser(user); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*user, true)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	err := user.HashPassword()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = user.CreateUser()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
