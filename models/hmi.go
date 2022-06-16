package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type HMI struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	MachineID uint           `json:"machine_id"`
	Name      string         `json:"name"`
	IP        pgtype.Inet    `json:"ip" gorm:"type:inet"`
	Port      int            `json:"port"`
	MAC       pgtype.Macaddr `json:"mac" gorm:"type:macaddr"`
	HMIUsers  []HMIUser      `json:"hmiusers" gorm:"foreignKey:ID"`
}
