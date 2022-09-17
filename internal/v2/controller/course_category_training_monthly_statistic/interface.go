package course_category_training_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSCategoryTrainingMonthlyStatistic(ctx *gin.Context)
}