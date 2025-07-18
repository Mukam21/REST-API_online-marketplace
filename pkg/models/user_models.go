package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string  `gorm:"unique;not null" json:"login"`
	Password string  `gorm:"not null" json:"password"`
	Orders   []Order `json:"orders,omitempty"`
}
