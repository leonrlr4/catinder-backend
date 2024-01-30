package handler

import (
	"catinder/internal/middleware"
	"net/http"

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
		user.POST("/register", RegisterUserHandler)

		user.GET("/profile", middleware.AuthMiddleware(), GetUserHandler)
	}

	// auth
	auth := r.Group("v1/auth")
	{
		auth.POST("/login", LocalLoginHandler)
		auth.GET("/google/login", GoogleLoginHandler)
		auth.GET("/google/callback", GoogleCallbackHandler)
	}

}
