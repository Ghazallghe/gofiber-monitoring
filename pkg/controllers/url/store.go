package url

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils/authService"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const UniqueValidationError = "23505"

func Store(c *fiber.Ctx) error {
	type inputData struct {
		Url       string `json:"url" validate:"required"`
		Threshold int32  `json:"threshold" validate:"required,min=1,max=512"`
	}

	var count int64
	db.DB.Model(&models.Url{}).Where("user_id = ?", authService.Id(c)).Count(&count)

	if count >= 20 {
		status := fiber.StatusBadRequest
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, "You reached the maximum urls."))
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

	userId, err := uuid.Parse(authService.Id(c))
	if err != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err.Error()))
	}

	url := models.Url{
		Url:       input.Url,
		Threshold: input.Threshold,
		UserId:    userId,
	}

	result := db.DB.Create(&url)

	if result.Error != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(url)
}
