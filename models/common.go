package models

import "time"

type Entity struct {
	ID        *string    `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"-"`
	IsActive  *bool      `json:"is_active,omitempty" gorm:"default:true"`
}

type Operator string

const (
	ADD      = "add"
	Multiply = "multiply"
)
