package handler

import (
	"catinder/internal/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you are in! ",
	})
}

func SetupRoutes(r *gin.Engine) {
	// home
	home := r.Group("/v1")
	{
		home.GET("/", HomePage)
	}

	// user
	user := r.Group("/v1/user")
	{
		user.GET("/:userId", GetUserHandler)
		user.POST("/register", RegisterUserHandler)
		user.POST("/login", LocalLoginHandler)
	}

	// google oauth
	auth := r.Group("/auth")
	{
		auth.GET("/google/login", GoogleLoginHandler)
		auth.GET("/google/callback", GoogleCallbackHandler)
	}

	// jwt
	r.Group("/v1", middleware.Auth(os.Getenv("JWT_SECRET")))
	{

	}
}
