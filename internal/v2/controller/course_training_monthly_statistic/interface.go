package course_training_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCourseTrainingMonthlyStatistic(ctx *gin.Context)
}

