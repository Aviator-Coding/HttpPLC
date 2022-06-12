package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     DBConfig
	Server ConfigServer
	JWT    JWTToken
}

type DBConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type ConfigServer struct {
	Url            string
	ReadTimeout    int
	TlsEnable      bool
	TlsCertificate string
	TlsKeyFile     string
}

type JWTToken struct {
	KeySecret        string
	KeyExpireMinutes int
}

var CFG *Config = LoadConfig()

// Load the Current Configuration
func LoadConfig() *Config {
	// Load Enviroment Variabels
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ENV]Error loading .env file")
	}

	cfg := &Config{
		DB: DBConfig{
			Host:     getStrEnv(os.Getenv("DB_HOST")),
			Port:     getStrEnv(os.Getenv("DB_PORT")),
			Password: getStrEnv(os.Getenv("DB_PASS")),
			User:     getStrEnv(os.Getenv("DB_USER")),
			SSLMode:  getStrEnv(os.Getenv("DB_SSLMODE")),
			DBName:   getStrEnv(os.Getenv("DB_NAME")),
		},
		Server: ConfigServer{
			Url:            getStrEnv(os.Getenv("SERVER_URL")),
			ReadTimeout:    getIntEnv(os.Getenv("SERVER_READ_TIMEOUT")),
			TlsEnable:      getBoolEnv(os.Getenv("SERVER_TLS_ENABLE")),
			TlsCertificate: getStrEnv(os.Getenv("SERVER_TLS_CERTIFICATE")),
			TlsKeyFile:     getStrEnv(os.Getenv("SERVER_TLS_KEY")),
		},
		JWT: JWTToken{
			KeySecret:        getStrEnv(os.Getenv("JWT_SECRET_KEY")),
			KeyExpireMinutes: getIntEnv(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")),
		},
	}
	return cfg
}

// Get the Enviroment Variable `Key` as a String
func getStrEnv(val string) string {
	if val == "" {
		log.Fatalf("[ENV]:Cannot Find Key:%v in the Enviroment Variables", val)
	}
	return val
}

// Get the Enviroment Variable `Key` as a Int
func getIntEnv(val string) int {
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("[ENV]:Cannot convert Key:%v to Integer", val)
	}
	return ret
}

// Get the Enviroment Variable `Key` as a Boolean
func getBoolEnv(val string) bool {
	ret, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("[ENV]:Cannot convert Key:%v to Boolean", val)
	}
	return ret
}
