package models

import "time"

type Line struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Machines  []Machine `json:"machines" gorm:"foreignKey:ID"`
}
