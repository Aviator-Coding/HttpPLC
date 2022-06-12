package models

import "gorm.io/gorm"

// Employe Database
type Employe struct {
	gorm.Model
	FirstName    string
	LastName     string
	Description  string
	BatchID      uint
	DepartmentID uint
	Violations   []Violation
	IsActive     bool
}

// Department Database
type Department struct {
	gorm.Model
	Name string
}

// Violations
type Violation struct {
	gorm.Model
	Text string
}
