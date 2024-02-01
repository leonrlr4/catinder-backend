package middleware

import (
	"github.com/gin-gonic/gin"
)

func LogoutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.Next()
	}
}
