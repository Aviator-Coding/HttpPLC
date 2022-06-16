package models

import "time"

type Machine struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LineID    uint      `json:"line_id"`
	Name      string    `json:"name"`
	Stations  []Station `json:"stastions" gorm:"foreignKey:ID"`
	PLCs      []PLC     `json:"plcs" gorm:"foreignKey:ID"`
	HMIs      []HMI     `json:"hmis" gorm:"foreignKey:ID"`
	IsActive  bool      `json:"is_active"`
}
