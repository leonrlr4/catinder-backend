package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you are in! mother fuxxer! ",
	})
}

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", HomePage)

		v1.GET("/user/:userId", GetUserHandler)
		v1.POST("/user/register", RegisterUserHandler)
	}
}
