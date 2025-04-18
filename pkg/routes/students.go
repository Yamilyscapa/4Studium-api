package std

import (
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of students",
	})
}
func GetStudentByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Student details",
	})
}
