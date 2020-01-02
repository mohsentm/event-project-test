package model

import (
	"time"
)

type Event struct {
	ID          int `gorm:"primary_key"json:"id"`
	Title       string `gorm:"type:varchar(100);not null,default:null",json:"title"`
	Description string `gorm:"type:text",json:"description"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}