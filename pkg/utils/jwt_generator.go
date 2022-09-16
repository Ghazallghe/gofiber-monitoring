package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

func JwtGenerator(user models.User) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	expTimeStr := os.Getenv("JWT_KEY_EXPIRE_HOUR")
	expTime, _ := strconv.Atoi(expTimeStr)
	exp := time.Hour * time.Duration(expTime)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(exp).Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(jwtKey)
	return t, err

}
