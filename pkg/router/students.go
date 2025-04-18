package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterStudentsRoute(r *gin.RouterGroup) {
	r.GET("/all", getAllStudents)
}

func getAllStudents(c *gin.Context) {

}
