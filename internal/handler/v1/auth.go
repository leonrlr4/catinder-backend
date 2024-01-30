package handler

import (
	"catinder/internal/dto"
	"catinder/internal/service"
	"catinder/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "514146514235-qfcu7tq5mjih5sh97vno4lj7taj46v4d.apps.googleusercontent.com",
	ClientSecret: "AIzaSyAF3Gl3j_4oRKFlwIzXD3jO2qZFC1FBmcM",
	RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
}

func GoogleLoginHandler(c *gin.Context) {
	goc := googleOauthConfig
	url := goc.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	code := c.Query("code")
	goc := googleOauthConfig
	token, err := goc.Exchange(context.Background(), code)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Exchange code"})
		return
	}

	// Use token to get user info from Google.
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer response.Body.Close()
	// Decode user info into a struct.
	userInfo := struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
		Picture  string `json:"picture"`
	}{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user info"})
		return
	}

	// create a user in your DB if it doesn't exist
	foundUser, err := service.GetUserByEmail(userInfo.Email)
	if err != nil {
		// create a new user
		_, err := service.RegisterUser(userInfo.UserName, userInfo.Email, userInfo.ID)
		if err != nil {
			util.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		// Generate the JWT token for the existing user
		token, err := util.GenerateToken(int(foundUser.ID))
		if err != nil {
			util.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return

	}

	// update the existing user
	foundUser.Username = userInfo.UserName
	foundUser.Picture = userInfo.Picture
	service.UpdateUser(foundUser)

	// Generate the JWT token for the new user
	newToken, err := util.GenerateToken(int(foundUser.ID))
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
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
	token, err := util.GenerateToken(int(user.ID))
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// https://accounts.google.com/o/oauth2/v2/auth?client_id=514146514235-qfcu7tq5mjih5sh97vno4lj7taj46v4d.apps.googleusercontent.com&redirect_uri=http://localhost:8080/v1/auth/google/callback&response_type=code&scope=profile%20email&state=state_string
