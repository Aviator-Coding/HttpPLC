package models

import "time"

type HMIUser struct {
	ID              uint      `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	HMIID           uint      `json:"hmi_id"`
	EmployeID       uint      `json:"employe_id"`
	PermissionLevel int       `json:"permission_level"`
}
