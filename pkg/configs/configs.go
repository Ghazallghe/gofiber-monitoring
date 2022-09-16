package configs

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

var JwtConfig func(c *fiber.Ctx) error

func SetUpConfigs() {
	setUpEnv()
	setUpJwt()
}

func setUpEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Couldn't load env")
	}
}

func setUpJwt() {
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtConfig = jwtware.New(jwtware.Config{
		SigningKey: jwtKey,
	})
}
