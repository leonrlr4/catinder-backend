package handler

import (
	"catinder/internal/dto"
	"catinder/internal/service"
	"catinder/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	// Bind JSON to dto.RegisterInfo struct
	var regInfo dto.RegisterInfo
	if err := c.ShouldBindJSON(&regInfo); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newUser, err := service.RegisterUser(regInfo)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"newUser": newUser})
}

func GetUserHandler(c *gin.Context) {
	// Get user ID from URL params
	userID := c.GetInt("userID")
	if userID == 0 {
		util.ErrorResponse(c, http.StatusBadRequest, "No user found")
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// IsLoggedInHandler checks if user is logged in
func IsLoggedInHandler(c *gin.Context) {
	userIDStr := c.Query("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}
	// check if user is logged in by checking if jwt token is empty
	user, err := service.GetUserByID(userID)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.JwtToken == "" {
		c.JSON(http.StatusOK, gin.H{"isLoggedIn": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"isLoggedIn": true})
}
