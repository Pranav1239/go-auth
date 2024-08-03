package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/profile", getProfile)
		user.PUT("/profile", updateProfile)
	}
}

func getProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get profile endpoint"})
}

func updateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update profile endpoint"})
}
