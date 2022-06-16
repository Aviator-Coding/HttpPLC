package models

import "time"

type Station struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	MachineID   uint      `json:"machine_id"`
	Name        string    `json:"name"`
	ShortName   string    `json:"short_name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}
