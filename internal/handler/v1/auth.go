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

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func getGoogleUserInfo(accessToken string) (GoogleUser, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return GoogleUser{
			ID:            "",
			Email:         "",
			VerifiedEmail: false,
			Picture:       "",
		}, err
	}
	defer response.Body.Close()
	// Decode the JSON response into our struct type.
	var userInfo GoogleUser
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		return GoogleUser{
			ID:            "",
			Email:         "",
			VerifiedEmail: false,
			Picture:       "",
		}, err
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
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, http.DefaultClient)

	token, err := goc.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	googleUserInfo, err := getGoogleUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	foundUser, err := service.GetUserByEmail(googleUserInfo.Email)
	if err != nil {
		regInfo.Email = googleUserInfo.Email
		regInfo.OAuthProvider = "google"
		regInfo.Picture = googleUserInfo.Picture
		regInfo.JwtToken = token.AccessToken
		regInfo.Username = googleUserInfo.ID
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

	foundUser.OAuthProvider = "google"
	foundUser.Picture = googleUserInfo.Picture
	if foundUser.Username == "" {
		foundUser.Username = googleUserInfo.ID
	}
	newToken, err := util.GenerateToken(int(foundUser.ID))
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	foundUser.JwtToken = newToken

	service.UpdateUser(foundUser)

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
