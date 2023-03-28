package user_unsubscribe_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSStatisticMonthlyUserUnsubscribe(ctx *gin.Context)
}
