package repository

import (
	"catinder/internal/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func CreateUser(user *model.User) error {
	result := db.Create(user)
	return result.Error
}

func InitializeDatabase(d *gorm.DB) {
	db = d
}
