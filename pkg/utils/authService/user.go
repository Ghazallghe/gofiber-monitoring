package authService

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx) (*models.User, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)

	dbUser := new(models.User)
	result := db.DB.Find(dbUser, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return dbUser, nil

}
