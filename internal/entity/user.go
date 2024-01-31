package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique"`
	Password       string `gorm:"type:varchar(255)"`
	Username       string `gorm:"type:varchar(255)"`
	Picture        string `gorm:"type:varchar(255)"`
	OAuthProvider  string `gorm:"type:varchar(100)"`
	JWTToken       string `gorm:"type:varchar(255)"`
	CreatedAt      string `gorm:"type:varchar(255)"`
	UpdatedAt      string `gorm:"type:varchar(255)"`
	gorm.DeletedAt `gorm:"index"`
}
