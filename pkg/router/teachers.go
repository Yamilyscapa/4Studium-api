package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterTeachersRoute(r *gin.RouterGroup) {
	r.GET("/all", getAllTeachers)
}

func getAllTeachers(c *gin.Context) {

}
