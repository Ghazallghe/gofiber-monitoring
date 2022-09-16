package url

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Index(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)

	dbUser := new(models.User)
	result := db.DB.Preload("Urls").Find(dbUser, "id = ?", id)
	if result.Error != nil {
		status := fiber.StatusUnauthorized
		return c.Status(status).JSON(utils.LogicalErrorHandling(status, result.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dbUser.Urls)
}
