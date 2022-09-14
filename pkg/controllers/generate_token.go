package controllers

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {
	type inputData struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	input := new(inputData)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := new(models.User)
	result := db.DB.Find(&user, "email = ?", input.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Email not found"})
	}

	if err := user.CheckPasswordHash(input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Email or Password is Not correct"})
	}

	t, err := utils.JwtGenerator(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t})
}
