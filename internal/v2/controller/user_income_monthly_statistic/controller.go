package user_income_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_income_monthly_statistic/api_get_trainer_income_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_income_monthly_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user_income_monthly_statistic.Resolver
}

func New(resolver user_income_monthly_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetTrainerIncomeMonthlyStatistic 獲取教練收益月數據
// @Summary 獲取教練收益月數據
// @Description 獲取教練收益月數據
// @Tags 教練課表數據_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_get_trainer_income_monthly_statistic.Output "Success!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/income_monthly_statistic [GET]
func (c *controller) GetTrainerIncomeMonthlyStatistic(ctx *gin.Context) {
	var input api_get_trainer_income_monthly_statistic.Input
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetTrainerIncomeMonthlyStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}
