package course

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCourses(ctx *gin.Context)
	GetCMSCourse(ctx *gin.Context)
	UpdateCMSCoursesStatus(ctx *gin.Context)
}
