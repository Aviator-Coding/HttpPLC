package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Aviator-Coding/HttpPLC/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect to the Database
var DB *gorm.DB 

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.CFG.DB.Host, configs.CFG.DB.Port, configs.CFG.DB.User, configs.CFG.DB.Password, configs.CFG.DB.DBName, configs.CFG.DB.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("[Database] - Unable to Connect %v", err)
	}

	log.Printf("[Database] - Connected to %s", &configs.CFG.DB.DBName)

	AutoMigrate(db)

	DB = db
}

// func NewConnection(config *configs.DBConfig) *gorm.DB {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
// 	)

// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold:             time.Second,   // Slow SQL threshold
// 			LogLevel:                  logger.Silent, // Log level
// 			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
// 			Colorful:                  false,         // Disable color
// 		},
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 	})

// 	if err != nil {
// 		log.Fatalf("[Database] - Unable to Connect %v", err)
// 	}

// 	log.Printf("[Database] - Connected to %s", config.DBName)

// 	AutoMigrate(db)
// 	return db
// }
