package repository

import (
	"catinder/internal/model"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func CreateUser(user *model.User) error {
	fmt.Println(user)
	result := db.Create(user)
	return result.Error
}

func InitializeDatabase(d *gorm.DB) {
	db = d
}
