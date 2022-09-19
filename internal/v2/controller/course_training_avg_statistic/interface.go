package course_training_avg_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCourseTrainingAvgStatistic(ctx *gin.Context)
}
