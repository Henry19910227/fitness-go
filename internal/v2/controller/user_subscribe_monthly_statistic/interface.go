package user_subscribe_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSUserSubscribeMonthlyStatistic(ctx *gin.Context)
}
