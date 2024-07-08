package models

import "time"

type Task struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Title     string     `json:"title" validate:"required,min=3,max=100" binding:"required"`
	Status    bool       `json:"status" validate:"boolean" binding:"required"`
}
