package url

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils/authService"
	"github.com/gofiber/fiber/v2"
)

const UniqueValidationError = "23505"

func Store(c *fiber.Ctx) error {
	type inputData struct {
		Url       string `json:"url" validate:"required"`
		Thershold int32  `json:"thershold" validate:"required,min=1,max=512"`
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

	user, err := authService.User(c)
	if err != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err.Error()))
	}

	url := models.Url{
		Url:       input.Url,
		Thershold: input.Thershold,
		UserId:    user.ID,
	}

	result := db.DB.Create(&url)

	if result.Error != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(url)
}
