package controllers

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(c *fiber.Ctx) error {
	type inputData struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	input := new(inputData)
	if err := c.BodyParser(input); err != nil {
		status := fiber.StatusBadRequest
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err.Error()))
	}

	if err := utils.ValidateStruct(input); err != nil {
		errType := fiber.ErrBadRequest
		return c.Status(errType.Code).JSON(utils.ValidatorErrorHandling(errType.Code, errType.Message, err))
	}

	user := new(models.User)
	result := db.DB.Find(&user, "email = ?", input.Email)
	if result.Error != nil {
		status := fiber.StatusBadRequest
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		status := fiber.StatusNotFound
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, "Email Not found"))
	}

	if err := user.CheckPasswordHash(input.Password); err != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, "Email or Password is incorrect"))
	}

	t, err := utils.JwtGenerator(*user)
	if err != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t})
}

func TestToken(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)

	dbUser := new(models.User)
	result := db.DB.Find(dbUser, "id = ?", id)
	if result.Error != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dbUser)
}
