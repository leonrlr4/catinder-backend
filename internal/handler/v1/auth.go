package handler

import (
	"catinder/internal/dto"
	"catinder/internal/service"
	"catinder/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	RedirectURL:  "YOUR_REDIRECT_URL",
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
	token, err := goc.Exchange(oauth2.NoContext, code)
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
		ID        string `json:"id"`
		Email     string `json:"email"`
		UserrName string `json:"user_name"`
		Picture   string `json:"picture"`
	}{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user info"})
		return
	}

	// Here you would create a user in your DB if it doesn't exist
	// or update the existing one, then generate a JWT.

	// Generate the JWT token for the user
	// ... (JWT token generation logic)

	c.JSON(http.StatusOK, gin.H{"token": "your_jwt_token_here"})
}

func LocalLoginHandler(c *gin.Context) {
	var loginInfo dto.LocalLoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// check is user  exist
	fmt.Println(loginInfo)

	user, err := service.GetUserByEmail(loginInfo.Email)
	// fmt.Println(user)
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
