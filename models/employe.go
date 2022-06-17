package models

import "time"

// Employe Database
type Employe struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	FirstName    string      `json:"firstname" validate:"required,min=2,max=30"`
	LastName     string      `json:"lastname" validate:"required,min=2,max=30"`
	Description  string      `json:"description,omitempty" validate:"omitempty,min=20,max=300"`
	BatchID      string      `json:"batch_id" gorm:"unique"`
	DepartmentID uint        `json:"department_id"`
	Violations   []Violation `json:"violations" gorm:"foreignKey:ID"`
	IsActive     bool        `json:"isactive"`
}
