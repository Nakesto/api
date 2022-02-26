package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	// apiRoutes.GET("/user/login", controller.UserLogin)
}
