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
	// home page route
	r.GET("/", HomePage)

	// user routes
	r.POST("/user/register", RegisterUser)
	r.GET("/user/:userId", GetUser)

}
