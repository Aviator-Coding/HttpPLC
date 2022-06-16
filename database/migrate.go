package database

import (
	"log"

	"github.com/Aviator-Coding/HttpPLC/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	log.Printf("[Database] - Start Automigrate %s", db.Name())

	log.Printf("[Database] - Migrate %s", "Department")
	err := db.AutoMigrate(&models.Department{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Department", err)
	}

	log.Printf("[Database] - Migrate %s", "User")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "User", err)
	}

	log.Printf("[Database] - Migrate %s", "Employe")
	err = db.AutoMigrate(&models.Employe{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Employe", err)
	}

	log.Printf("[Database] - Migrate %s", "HMI")
	err = db.AutoMigrate(&models.HMI{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "HMI", err)
	}

	log.Printf("[Database] - Migrate %s", "HMIUser")
	err = db.AutoMigrate(&models.HMIUser{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "HMIUser", err)
	}

	log.Printf("[Database] - Migrate %s", "Line")
	err = db.AutoMigrate(&models.Line{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Line", err)
	}

	log.Printf("[Database] - Migrate %s", "Machine")
	err = db.AutoMigrate(&models.Machine{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Machine", err)
	}

	log.Printf("[Database] - Migrate %s", "PLC")
	err = db.AutoMigrate(&models.PLC{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "PLC", err)
	}

	log.Printf("[Database] - Migrate %s", "Station")
	err = db.AutoMigrate(&models.Station{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Station", err)
	}

	log.Printf("[Database] - Migrate %s", "Violation")
	err = db.AutoMigrate(&models.Violation{})
	if err != nil {
		log.Fatalf("[Database] - Migrate %s - Error %s", "Violation", err)
	}

	log.Printf("[Database] - Done Automigrate %s", db.Name())

}
