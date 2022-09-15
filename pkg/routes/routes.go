package routes

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	r := app.Group("/api")

	// user
	{
		user := r.Group("/users")
		user.Post("", controllers.CreateUser)
		user.Post("/token", controllers.GenerateToken)
	}
}