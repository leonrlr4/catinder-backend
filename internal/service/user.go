// service/user.go

package service

import (
	"catinder/internal/dto"
	"catinder/internal/entity"
	"catinder/internal/repository"
	"catinder/util"
	"fmt"
)

func RegisterUser(registerInfo dto.RegisterInfo) (*entity.User, error) {
	isUserFound := repository.IsUserExist(registerInfo.Email)
	if isUserFound {
		return nil, fmt.Errorf("user already exists")
	}

	// Hash password
	hashedPassword, err := util.HashPassword(registerInfo.Password)
	if err != nil {
		return nil, err
	}

	// Create new user
	newUser := &entity.User{
		Username:      registerInfo.Username,
		Email:         registerInfo.Email,
		Password:      hashedPassword,
		OAuthProvider: registerInfo.OAuthProvider,
		Picture:       registerInfo.Picture,
		JwtToken:      registerInfo.JwtToken,
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

func GetUserByEmail(email string) (*entity.User, error) {
	// Find user by email
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserFields(user *entity.User, updatedFields map[string]interface{}) error {
	err := repository.UpdateUserFields(user.ID, updatedFields)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *entity.User) error {
	return repository.UpdateUser(user)
}

func GenerateTokenAndUpdateUser(user *entity.User) (string, error) {
	token, err := util.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	user.JwtToken = token

	if err := repository.UpdateUser(user); err != nil {
		return "", fmt.Errorf("could not update user: %v", err)
	}

	return token, nil
}

func GetUserByID(userID int) (*entity.User, error) {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("user:", user)
	return user, nil
}
