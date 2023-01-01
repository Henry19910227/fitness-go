package user_course_usage_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetTrainerCourseUsageMonthlyStatistic(ctx *gin.Context)
	Statistic()
}
