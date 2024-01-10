package handler

import (
	"catinder/internal/model"
	"catinder/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// parse request body
type RegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// handle htte request
func RegisterUser(c *gin.Context) {
	var regInfo RegisterInfo
	if err := c.ShouldBindJSON(&regInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration information"})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := model.User{
		Username: regInfo.Username,
		Password: string(hashedPassword),
	}

	// call user repository to create user
	if err := repository.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
