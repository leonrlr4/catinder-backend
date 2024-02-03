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
	r.Use(middleware.CorsMiddleware())

	// home
	home := r.Group("/v1")
	{
		home.GET("/", HomePage)
	}

	// user
	user := r.Group("/v1/user")
	{
		user.POST("/register", RegisterUserHandler)
		user.POST("/logout", LogoutHandler)
		user.POST("/login", LocalLoginHandler)

		user.GET("/isLoggedIn", IsLoggedInHandler)
		user.GET("/profile", middleware.AuthMiddleware(), GetUserHandler)
	}

	// auth
	auth := r.Group("v1/auth")
	{
		auth.GET("/google/login", GoogleLoginHandler)
		auth.GET("/google/callback", GoogleCallbackHandler)
	}

}
