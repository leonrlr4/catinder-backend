package handler

import (
	"catinder/internal/dto"
	"catinder/internal/service"
	"catinder/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	// Bind JSON to dto.RegisterInfo struct
	var regInfo dto.RegisterInfo
	if err := c.ShouldBindJSON(&regInfo); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 調用 service 層的 RegisterUser 函數
	_, err := service.RegisterUser(regInfo.Username, regInfo.Email, regInfo.Password)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func GetUserHandler(c *gin.Context) {
	// Get user ID from URL params
	userID := c.Param("userId")
	if userID == "" {
		util.ErrorResponse(c, http.StatusBadRequest, "No user ID provided")
		return
	}

	// 調用 service 層的 GetUser 函數
	user, err := service.GetUser(userID)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
