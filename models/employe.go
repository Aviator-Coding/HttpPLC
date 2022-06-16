package models

import "time"

// Employe Database
type Employe struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	FirstName    string      `json:"firstname"`
	LastName     string      `json:"lastname"`
	Description  string      `json:"description"`
	BatchID      uint        `json:"batch_id"`
	DepartmentID uint        `json:"department_id"`
	Violations   []Violation `json:"violations" gorm:"foreignKey:ID"`
	IsActive     bool        `json:"isactive"`
}
