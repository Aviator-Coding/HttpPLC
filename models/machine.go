package models

import "gorm.io/gorm"

type Machine struct {
	gorm.Model
	Name     string
	Stations []Station
	PLCs     []PLC
	HMIs     []HMI
	IsActive bool
}

type Station struct {
	gorm.Model
	Name     string
	IsActive bool
}

type PLC struct {
	gorm.Model
	Name string
}

type HMI struct {
	gorm.Model
	Name  string
	Users []User
}
