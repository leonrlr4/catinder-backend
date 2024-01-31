// service/user.go

package service

import (
	"catinder/internal/entity"
	"catinder/internal/repository"
	"catinder/util"
	"fmt"
)

// RegisterUser 處理註冊的業務邏輯
func RegisterUser(username, email, password, OAuthProvider string) (*entity.User, error) {
	// Check if user exists
	isUserFound := repository.IsUserExist(email)
	if isUserFound {
		return nil, fmt.Errorf("user already exists")
	}

	// Hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create new user
	newUser := &entity.User{
		Username:      username,
		Email:         email,
		Password:      hashedPassword,
		OAuthProvider: OAuthProvider,
		Picture:       "",
		JWTToken:      "",
		CreatedAt:     util.GetCurrentTime().String(),
		UpdatedAt:     util.GetCurrentTime().String(),
	}

	// Create user
	repository.CreateUser(newUser)
	_, err = GenerateTokenAndUpdateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUserByID
func GetUser(userID int) (*entity.User, error) {
	// Find user by ID
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail
func GetUserByEmail(email string) (*entity.User, error) {
	// Find user by email
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser
func UpdateUser(user *entity.User) error {
	return repository.UpdateUser(user)
}

func GenerateTokenAndUpdateUser(user *entity.User) (string, error) {
	token, err := util.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	user.JWTToken = token

	if err := repository.UpdateUser(user); err != nil {
		return "", fmt.Errorf("could not update user: %v", err)
	}

	return token, nil
}
