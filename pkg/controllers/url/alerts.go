package url

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Alerts(c *fiber.Ctx) error {
	urlId := c.Params("id")

	status, err := checkPermission(c)
	if status != fiber.StatusOK {
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err))
	}

	alerts := new([]models.Alert)
	result := db.DB.Find(alerts, "url_id = ?", urlId)

	if result.Error != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(alerts)
}
