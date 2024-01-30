package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CorsMiddleware() gin.HandlerFunc {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // 允許的來源
		AllowedMethods:   []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
		AllowCredentials: true,
	})

	return func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
