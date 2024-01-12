// service/user.go

package service

import (
	"catinder/internal/entity"
	"catinder/internal/repository"
	"catinder/util"
)

// RegisterUser 處理註冊的業務邏輯，不涉及 HTTP 層面的操作。
func RegisterUser(username, email, password string) (*entity.User, error) {
	// Hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := entity.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	// Call user repository to create user
	if err := repository.CreateUser(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}

// GetUser 處理獲取用戶的業務邏輯，不涉及 HTTP 層面的操作。
func GetUser(userID string) (*entity.User, error) {
	// Find user by ID
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
