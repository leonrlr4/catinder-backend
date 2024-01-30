// service/user.go

package service

import (
	"catinder/internal/entity"
	"catinder/internal/repository"
	"catinder/util"
)

type newUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
}

// RegisterUser 處理註冊的業務邏輯
func RegisterUser(username, email, password string) (*entity.User, error) {
	hashedPassword := ""
	if password != "" {
		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			return nil, err
		}
		newUser := entity.User{
			Username: username,
			Email:    email,
			Password: hashedPassword,
			Picture:  "",
		}
	}
	// Hash password
	newUser := entity.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Picture:  "",
	}

	// Call user repository to create user
	if err := repository.CreateUser(&newUser); err != nil {
		return nil, err
	}
	return &newUser, nil
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
