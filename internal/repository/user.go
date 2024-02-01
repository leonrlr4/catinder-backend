package repository

import (
	"catinder/internal/entity"
	"fmt"

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
		Select("ID, Username, Email, Picture, jwt_token").
		Omit("Password", "DeletedAt").
		First(&user, userID)

	fmt.Println(&user)
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
func UpdateUserFields(id uint, updatedFields map[string]interface{}) error {
	fmt.Println("UpdateUserFields", updatedFields)
	// Start a new transaction
	tx := db.Begin()

	// Update only the specified fields
	tx = tx.Model(&entity.User{}).Where("id = ?", id).Updates(updatedFields)

	// Commit the transaction
	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
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
