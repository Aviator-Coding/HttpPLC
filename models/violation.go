package models

import "time"

// Violations
type Violation struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text" validate:"required,min=20,max=300"`
}
