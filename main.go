package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/join-api/routes"
)

func main() {

	// Initialize the database
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	publicRoutes := r.Group("/")
	{
		publicRoutes.GET("aku", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "aku",
			})
		})
	}

	apiRoutes := r.Group("/auth")
	{
		apiRoutes.Use(func(c *gin.Context) {
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
			} else if authHeader != "Bearer 12345" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
			} else {
				if len(c.Keys) == 0 {
					c.Keys = make(map[string]interface{})
				}
				fmt.Print(c.Keys)
				fmt.Print("\n")
				c.Keys["auth"] = authHeader
				fmt.Print(c.Keys)
				fmt.Print("\n")
				c.Next()
			}
		})
		routes.UserRoutes(apiRoutes)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
