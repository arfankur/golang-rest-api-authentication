package main

import (
	"go-auth/config"
	"go-auth/controllers"
	"go-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db := config.InitDB()
	defer config.CloseDB(db)

	// Initialize Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected route
	r.GET("/profile", middlewares.Authenticate, controllers.Profile)

	// Start the server
	r.Run(":8080")
}
