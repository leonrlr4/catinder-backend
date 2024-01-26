package entity

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Picture  string `gorm:"type:varchar(255);"`
	Password string `gorm:"not null"`
}
