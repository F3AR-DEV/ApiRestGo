package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"Description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
