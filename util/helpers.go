package util

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// generate jwt token
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	// Set expiration time to 1 day from now
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now().Unix(),
		"exp":     expirationTime.Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// parse jwt token
func ParseToken(tokenString string) (int, error) {
	tokenString = strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return int(claims["user_id"].(float64)), nil
}

type Endpoint struct {
	AuthURL  string `json:"auth_url"`
	TokenURL string `json:"token_url"`
}
type GoogleUserInfo struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url"`
	Scopes       []string `json:"scopes"`
	Endpoint     Endpoint `json:"endpoint"`
}

// get googleOauthConfig from env
func GetGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

// GetCurrentTime
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}
