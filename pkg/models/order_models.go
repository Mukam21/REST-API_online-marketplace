package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:500;not null" json:"description"`
	ImageURL    string `gorm:"not null" json:"image_url"`
	Price       uint   `gorm:"not null" json:"price"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user"`
}
