package user_register_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSUserRegisterMonthlyStatistic(ctx *gin.Context)
}
