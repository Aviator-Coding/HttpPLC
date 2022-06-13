package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Line struct {
	gorm.Model
	Name     string
	Machines []Machine
}

type Machine struct {
	gorm.Model
	LineID   uint
	Name     string
	Stations []Station
	PLCs     []PLC
	HMIs     []HMI
	IsActive bool
}

type Station struct {
	gorm.Model
	MachineID   uint
	Name        string
	ShortName   string
	Description string
	IsActive    bool
}

type PLC struct {
	gorm.Model
	MachineID uint
	Name      string
	IP        pgtype.Inet `gorm:"type:inet"`
	Port      int
	MAC       pgtype.Macaddr `gorm:"type:macaddr"`
}

type HMI struct {
	gorm.Model
	MachineID uint
	Name      string
	IP        pgtype.Inet `gorm:"type:inet"`
	Port      int
	MAC       pgtype.Macaddr `gorm:"type:macaddr"`
	HMIUsers  []HMIUser
}

type HMIUser struct {
	gorm.Model
	HMIID           uint
	EmployeID       uint
	PermissionLevel int
}
