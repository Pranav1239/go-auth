package main

import (
	routes "Project/router"
	utils "Project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.SetupRoutes(router)
	utils.SetTrustedProxies(router)
	utils.SetupCORS(router)
	router.NoRoute(utils.HandleNotFound)

	port := utils.GetPort()
	fmt.Printf("Listening to %v", port)
	router.Run(":" + port)
}
