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

		userID, _ := util.ParseToken(token)
		if userID == 0 {
			util.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Check if the token is still valid for the user
		user, _ := service.GetUserByID(userID)
		if ("Bearer " + user.JwtToken) != token {
			util.ErrorResponse(c, http.StatusUnauthorized, "Invalid2 token")
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
