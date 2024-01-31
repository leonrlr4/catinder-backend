package repository

import (
	"catinder/internal/entity"

	"gorm.io/gorm"
)

var db *gorm.DB

func CreateUser(user *entity.User) (int, error) {
	result := db.Create(user)
	return int(user.ID), result.Error

}

func FindUserByID(userID int) (*entity.User, error) {
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

// UpdateUser
func UpdateUser(user *entity.User) error {
	result := db.Save(user)
	return result.Error
}

// check if user exists
func IsUserExist(email string) bool {
	var user entity.User
	result := db.
		Select("ID").
		Where("email = ?", email).
		First(&user)

	return result.Error == nil
}
