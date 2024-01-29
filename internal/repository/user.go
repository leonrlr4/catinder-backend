package repository

import (
	"catinder/internal/entity"

	"gorm.io/gorm"
)

var db *gorm.DB

func CreateUser(user *entity.User) error {
	result := db.Create(user)
	return result.Error
}

func FindUserByID(userID string) (*entity.User, error) {
	var user entity.User
	result := db.
		Select("ID, Username, Email, Picture").
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

func FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := db.
		Select("ID, Username, Email, Password, Picture").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
