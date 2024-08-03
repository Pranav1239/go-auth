package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}

func register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register endpoint"})
}
