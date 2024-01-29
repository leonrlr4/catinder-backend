package util

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// generate jwt token
func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
