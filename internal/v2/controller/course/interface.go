package course

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetFavoriteCourses(ctx *gin.Context)
	GetCMSCourses(ctx *gin.Context)
	GetCMSCourse(ctx *gin.Context)
	UpdateCMSCoursesStatus(ctx *gin.Context)
	UpdateCMSCoursesCover(ctx *gin.Context)
	CreatePersonalCourse(ctx *gin.Context)
}
