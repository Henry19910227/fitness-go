package course_release_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCourseReleaseMonthlyStatistic(ctx *gin.Context)
}
