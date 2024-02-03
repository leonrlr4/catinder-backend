package handler

import (
	"catinder/internal/dto"
	"catinder/internal/service"
	"catinder/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func GoogleLoginHandler(c *gin.Context) {
	goc := util.GetGoogleOauthConfig()
	url := goc.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func getUserInfo(accessToken string) (*dto.UserInfo, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	userInfo := &dto.UserInfo{}
	err = json.NewDecoder(response.Body).Decode(userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func LocalLoginHandler(c *gin.Context) {
	var loginInfo dto.LocalLoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := service.GetUserByEmail(loginInfo.Email)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "user not exist")
		return
	}

	if !util.CheckPasswordHash(loginInfo.Password, user.Password) {
		util.ErrorResponse(c, http.StatusBadRequest, "wrong password")
		return
	}

	_, err = service.GenerateTokenAndUpdateUser(user)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser := dto.ReturnLoginInfo{
		ID:            int(user.ID),
		UserName:      user.Username,
		Email:         user.Email,
		Picture:       user.Picture,
		OAuthProvider: user.OAuthProvider,
		JwtToken:      user.JwtToken,
		CreatedAt:     user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func LogoutHandler(c *gin.Context) {
	userIDStr := c.Query("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("Error parsing user ID:", err)
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	updatedFields := map[string]interface{}{
		"jwt_token": "",
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Failed to get user")
		return
	}

	service.UpdateUserFields(user, updatedFields)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func GoogleCallbackHandler(c *gin.Context) {
	var regInfo dto.RegisterInfo

	code := c.Query("code")
	goc := util.GetGoogleOauthConfig()

	token, err := goc.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code"})
		return
	}

	userInfo, err := getUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	foundUser, err := service.GetUserByEmail(userInfo.Email)
	if err != nil {
		newUser, err := service.RegisterUser(regInfo)
		if err != nil {
			util.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		token, err := util.GenerateToken(int(newUser.ID))
		if err != nil {
			util.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	if foundUser.OAuthProvider == "" {
		foundUser.OAuthProvider = "google"
		foundUser.JwtToken = token.AccessToken
		service.UpdateUser(foundUser)
	}

	foundUser.Username = userInfo.Username
	foundUser.Picture = userInfo.Picture
	service.UpdateUser(foundUser)

	newToken, err := util.GenerateToken(int(foundUser.ID))
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
