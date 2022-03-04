package models

import "time"

type Example struct {
	Data      string     `json:"data"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

// TableName of this model
func (e *Example) TableName() string {
	return "examples"
}

// create db table like auto migration
type examples interface {
	TableName() string
}
