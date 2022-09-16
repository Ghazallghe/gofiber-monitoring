package url

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils/authService"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	id := authService.Id(c)

	dbUser := new(models.User)
	result := db.DB.Preload("Urls").Find(dbUser, "id = ?", id)
	if result.Error != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dbUser.Urls)
}
