package configs

import "github.com/joho/godotenv"

func SetUpEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Couldn't load env")
	}
}
