package handler

import (
	"catinder/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you are in! mother fuxxer! ",
	})
}

func SetupRoutes(r *gin.Engine) {
	// v1 API routes
	v1 := r.Group("/v1")
	{
		v1.GET("/", HomePage)
		v1.POST("/user/register", service.RegisterUser)
		v1.GET("/user/:userId", service.GetUser)
	}
}
