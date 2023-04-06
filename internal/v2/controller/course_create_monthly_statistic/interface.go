package course_create_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSStatisticMonthlyCourseCreate(ctx *gin.Context)

	Statistic()
}
