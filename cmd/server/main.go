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
	r := gin.Default()
	r.Use(cors.Default())

	Init(r)
}

func Init(r *gin.Engine) {
	Config() // Initialize environment variables

	Router(r) // Initialize routes

	// Initialize the server
	if PORT := os.Getenv("PORT"); PORT != "" {
		r.Run(":" + PORT)
	} else {
		r.Run(":8080")
	}
}

func Config() {
	// Load .env file
	godotenv.Load()
}

func Router(r *gin.Engine) {

	// Initialize routes
	students := r.Group("/students")
	teachers := r.Group("/teachers")

	// Register routes
	router.RegisterStudentsRoute(students)
	router.RegisterTeachersRoute(teachers)
}
