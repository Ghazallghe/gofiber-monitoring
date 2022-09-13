package main

import (
	"github.com/Ghazallghe/gofiber-monitoring/pkg/configs"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	initialization()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

func initialization() {
	configs.SetUpEnv()
	db.SetUpDB()
}
