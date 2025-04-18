package main

import (
	"os"

	"studium/pkg/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Studium API
// @version 0.1
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

	// Register routes
	students := r.Group("/students")
	teachers := r.Group("/teachers")
	router.RegisterStudentsRoute(students)
	router.RegisterTeachersRoute(teachers)

	// Initialize the server
	if PORT := os.Getenv("PORT"); PORT != "" {
		r.Run(":" + PORT)
	} else {
		r.Run(":8080")
	}
}
