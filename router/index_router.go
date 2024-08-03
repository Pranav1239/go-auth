package routes

import (
	"Project/router/auth"
	"Project/router/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1")

	auth.SetupAuthRoutes(api)
	user.SetupUserRoutes(api)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})
}
