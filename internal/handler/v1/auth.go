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

func GoogleCallbackHandler(c *gin.Context) {
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
		newUser, err := service.RegisterUser(userInfo.UserName, userInfo.Email, userInfo.PassWord, "google")
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

	foundUser.Username = userInfo.UserName
	foundUser.Picture = userInfo.Picture
	service.UpdateUser(foundUser)

	newToken, err := util.GenerateToken(int(foundUser.ID))
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

func getUserInfo(accessToken string) (*struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	UserName      string `json:"user_name"`
	PassWord      string `json:"password"`
	Picture       string `json:"picture"`
	OAuthProvider string `json:"o_auth_provider"`
	JwtToken      string `json:"jwt_token"`
}, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	userInfo := &struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		UserName      string `json:"user_name"`
		PassWord      string `json:"password"`
		Picture       string `json:"picture"`
		OAuthProvider string `json:"o_auth_provider"`
		JwtToken      string `json:"jwt_token"`
	}{}

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

	// check is user  exist
	user, err := service.GetUserByEmail(loginInfo.Email)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "user not exist")
		return
	}

	// check password
	if !(util.CheckPasswordHash(loginInfo.Password, user.Password)) {
		util.ErrorResponse(c, http.StatusBadRequest, "wrong password")
		return
	}

	// Generate the JWT token for the user
	token, err := service.GenerateTokenAndUpdateUser(user)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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
		fmt.Println("Error getting user:", err)
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = service.UpdateUserFields(user, updatedFields)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
