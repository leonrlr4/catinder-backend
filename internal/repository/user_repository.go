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

func FindUserByID(userID string) (*model.User, error) {
	var user model.User
	result := db.
		Select("ID, Username, Email").
		Omit("Password", "DeletedAt").
		First(&user, userID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func InitializeDatabase(d *gorm.DB) {
	db = d
}
