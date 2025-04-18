package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	Init()
}

func Init() {
	// Load .env file
	godotenv.Load()

	// Initialize the Gin router
	r := gin.Default()
	// Set up CORS middleware
	r.Use(cors.Default())

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to 4Studium API",
		})
	})

	// Initialize the server
	if PORT := os.Getenv("PORT"); PORT != "" {
		r.Run(":" + PORT)
	} else {
		r.Run(":8080")
	}
}
