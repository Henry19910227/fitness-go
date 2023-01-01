package user_income_monthly_statistic

import "github.com/gin-gonic/gin"

type Controller interface {
	GetTrainerIncomeMonthlyStatistic(ctx *gin.Context)
	Statistic()
}
