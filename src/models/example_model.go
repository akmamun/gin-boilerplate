package models

import "time"

type Example struct {
	ID        int
	Data      string     `json:"data"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

// Database TableName of this model

func (e *Example) TableName() string {
	return "examples"
}
