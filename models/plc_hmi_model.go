package models

import (
	"time"
)

type PlcHMIUser struct {
	FirstName string    `json:"firstname,omitempty" validate:"required"`
	LastName  string    `json:"lastname,omitempty" validate:"required"`
	BatchID   string    `json:"batchid,omitempty" validate:"required"`
	CreatedAt time.Time `json:"createdat,omitempty"`
	UpdatedAt time.Time `json:"updatedat,omitempty"`
	Location  string    `json:"location,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Machines  []PLCMachines
}

type PLCMachines struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Stations []PLCMachineStations
}

type PLCMachineStations struct {
	Name         string      `json:"name,omitempty" validate:"required"`
	Location     string      `json:"location,omitempty" validate:"required"`
	Permission   bool        `json:"permission,omitempty" validate:"required"`
	LastLogin    time.Time   `json:"lastlogin,omitempty"`
	LoginHistory []time.Time `json:"loginHistory,omitempty"`
}
