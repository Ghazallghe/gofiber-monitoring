package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpDB() {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" TimeZone=" + os.Getenv("SERVER_TIMEZONE")

	retries := 5
	var err error

	for retries > 0 {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Couldn't connect to database")
		retries--
		time.Sleep(5 * time.Second)
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
}
