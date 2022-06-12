package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Get the Enviroment Variable `Key` as a String
func getStrEnv(key string) string {
	// Load the Enviroment Variables first
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ENV]Error loading .env file")
	}

	// Get the Enviroment Variables
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("[ENV]:Cannot Find Key:%v in the Enviroment Variables", key)
	}
	return val
}

// Get the Enviroment Variable `Key` as a Int
func getIntEnv(key string) int {
	val := getStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("[ENV]:Cannot convert Key:%v to Integer", key)
	}
	return ret
}

// Get the Enviroment Variable `Key` as a Boolean
func getBoolEnv(key string) bool {
	val := getStrEnv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("[ENV]:Cannot convert Key:%v to Boolean", key)
	}
	return ret
}

func EnvMongoURI() string {
	return getStrEnv("MONGOURI")
}

func EnvMongoDatabase() string {
	return getStrEnv("MONGODATABASE")
}

func EnvServerUrl() string {
	return getStrEnv("SERVER_URL")
}

func EnvServerTlsEnable() bool {
	return getBoolEnv("SERVER_TLS_ENABLE")
}

func EnvServerTlsCertificate() string {
	return getStrEnv("SERVER_TLS_CERTIFICATE")
}

func EnvServerTlsKey() string {
	return getStrEnv("SERVER_TLS_KEY")
}

func EnvJwtSecretKey() string {
	return getStrEnv("JWT_SECRET_KEY")
}

func EnvJwtSecretKeyExpireMinutesCount() int {
	return getIntEnv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
}
