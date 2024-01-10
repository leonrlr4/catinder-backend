package handler

import (
	"catinder/internal/model"
	"catinder/internal/repository"
	"fmt"
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
	fmt.Println(c)
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

func GetUser(c *gin.Context) {
	// get user ID from URL params
	userID := c.Param("userId")
	fmt.Println(userID)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user ID provided"})
		return
	}

	// find user by ID
	user, err := repository.FindUserByID(userID)
	if err != nil {
		// error handling
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find user"})
		return
	}

	// return found user
	c.JSON(http.StatusOK, gin.H{"user": user})
}
