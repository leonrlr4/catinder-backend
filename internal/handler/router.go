package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you are in! mother fucker! ",
	})
}

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", RegisterUser)
	r.GET("/", HomePage)
}
