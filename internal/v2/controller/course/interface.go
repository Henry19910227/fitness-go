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

	CreateUserCourse(ctx *gin.Context)
	GetUserCourses(ctx *gin.Context)
	DeleteUserCourse(ctx *gin.Context)
	UpdateUserCourse(ctx *gin.Context)
	GetUserCourse(ctx *gin.Context)
	GetUserCourseStructure(ctx *gin.Context)

	GetTrainerCourses(ctx *gin.Context)
	GetTrainerCourseOverview(ctx *gin.Context)
	CreateTrainerCourse(ctx *gin.Context)
	GetTrainerCourse(ctx *gin.Context)
	UpdateTrainerCourse(ctx *gin.Context)
	DeleteTrainerCourse(ctx *gin.Context)
	SubmitTrainerCourse(ctx *gin.Context)

	GetStoreCourse(ctx *gin.Context)
	GetStoreCourses(ctx *gin.Context)
	GetStoreCourseStructure(ctx *gin.Context)
	GetStoreTrainerCourses(ctx *gin.Context)
	GetStoreHomePage(ctx *gin.Context)
}
