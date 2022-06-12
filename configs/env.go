package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvServerUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SERVER_URL")
}

func EnvJwtSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("JWT_SECRET_KEY")
}

func EnvJwtSecretKeyExpireMinutesCount() int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	minutesCount, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil {
		log.Fatal("Error cannot Convert JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT to Int")
	}
	return minutesCount
}
