// middleware/auth_middleware.go

package middleware

import (
	"catinder/internal/service"
	"catinder/util"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.ErrorResponse(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		userID, err := util.ParseToken(token)

		if err != nil {
			util.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Check if the token is still valid for the user
		user, err := service.GetUserByID(userID)
		if err != nil || user.JwtToken != token {
			util.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
