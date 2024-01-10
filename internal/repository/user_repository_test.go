package repository

import (
	"catinder/internal/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Create a mock user
	user := &model.User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Call the CreateUser function
	err := CreateUser(user)

	// Check if there was an error
	assert.NoError(t, err)

	fmt.Println(user)
}
