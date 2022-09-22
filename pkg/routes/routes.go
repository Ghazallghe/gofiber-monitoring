package routes

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/configs"
	urlHandler "github.com/Ghazallghe/gofiber-monitoring/pkg/controllers/url"
	userHandler "github.com/Ghazallghe/gofiber-monitoring/pkg/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	r := app.Group("/api")
	// user
	{
		user := r.Group("/users")
		user.Post("", userHandler.CreateUser)
		user.Post("/token", userHandler.GenerateToken)
		// test token
		user.Use(configs.JwtConfig)
		user.Get("/me", userHandler.TestToken)
	}

	// url
	{
		url := r.Group("/urls")
		url.Use(configs.JwtConfig)

		url.Get("", urlHandler.Index)
		url.Post("", urlHandler.Store)
		url.Get("/:id/statistics", urlHandler.Statistics)
		url.Get("/:id/alerts", urlHandler.Alerts)
	}
}
