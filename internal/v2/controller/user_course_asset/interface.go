package user_course_asset

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateCMSCourseUsers(ctx *gin.Context)
}
