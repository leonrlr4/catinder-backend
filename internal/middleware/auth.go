// middleware/auth_middleware.go

package middleware

import (
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

		c.Set("userID", userID)
		c.Next()
	}
}

// func Auth(secretKey string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 检查Authorization header是否存在
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
// 			return
// 		}

// 		// 检查Bearer token格式
// 		bearerToken := strings.Split(authHeader, " ")
// 		if len(bearerToken) != 2 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 			return
// 		}

// 		// 解析token
// 		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method")
// 			}
// 			return []byte(secretKey), nil
// 		})

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			c.Set("userID", claims["user_id"])
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}
// 	}
