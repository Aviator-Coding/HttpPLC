package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to the Database
var DBConnection *gorm.DB = NewConnection(&CFG.DB)

func NewConnection(config *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Database] - Unable to Connect %v", err)
	}
	log.Printf("[Database] - Connected to %s", config.DBName)
	return db
}
