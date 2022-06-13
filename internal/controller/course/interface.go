package course

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCourses (ctx *gin.Context)
}
