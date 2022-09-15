package controllers

import (
	"errors"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

const UniqueValidationError = "23505"

func CreateUser(c *fiber.Ctx) error {
	type inputData struct {
		Email     string `json:"email" validate:"required,email"`
		FirstName string `json:"first_name" validate:"required,min=3,max=64"`
		LastName  string `json:"last_name" validate:"required,min=3,max=64"`
		Password  string `json:"password" validate:"required,min=8,max=64"`
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

	user := models.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
	}

	if err := user.HashPassword(); err != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err.Error()))
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		status := fiber.StatusBadRequest
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) && pgErr.Code == UniqueValidationError {
			return c.Status(status).JSON(utils.LogicalErrorHandling(status, "Email is already used"))
		}
		status = fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
