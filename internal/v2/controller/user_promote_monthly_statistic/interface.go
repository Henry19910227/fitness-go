package user_promote_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSUserPromoteMonthlyStatistic(ctx *gin.Context)
}
