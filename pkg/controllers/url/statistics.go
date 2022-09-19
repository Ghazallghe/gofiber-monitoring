package url

import (
	"time"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils/authService"
	"github.com/gofiber/fiber/v2"
)

func Statistics(c *fiber.Ctx) error {
	urlId := c.Params("id")
	date := c.Query("date")

	if date == "" {
		date = time.Now().Format("01-02-2006")
	}

	status, err := checkPermission(c)
	if status != fiber.StatusOK {
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, err))
	}

	stat := new(models.Statistics)
	result := db.DB.Find(stat, "url_id = ? AND date = ?", urlId, date)

	if result.Error != nil {
		status := fiber.StatusInternalServerError
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		status := fiber.StatusNotFound
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, "There's no statistics available for your request"))
	}

	totalSuccess := stat.StatusOk
	totalFailure := stat.ClientError + stat.ServerError

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"detail": stat, "total_success": totalSuccess, "total_failure": totalFailure})
}

func checkPermission(c *fiber.Ctx) (int, string) {
	urlId := c.Params("id")
	userId := authService.Id(c)

	url := new(models.Url)

	result := db.DB.Find(&url, "id = ?", urlId)

	if result.Error != nil {
		status := fiber.StatusInternalServerError
		return status, result.Error.Error()
	}

	if result.RowsAffected == 0 {
		status := fiber.StatusNotFound
		return status, "Url does Not exists"
	}

	if url.UserId.String() != userId {
		status := fiber.StatusForbidden
		return status, "You are Not allowed to observer this url"
	}

	return fiber.StatusOK, ""
}
