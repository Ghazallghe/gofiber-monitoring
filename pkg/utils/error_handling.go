package utils

import "github.com/gofiber/fiber/v2"

func ValidatorErrorHandling(status int, message string, errors []map[string]string) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
		"errors":  errors,
	}
}

func LogicalErrorHandling(status int, message string) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
	}
}
