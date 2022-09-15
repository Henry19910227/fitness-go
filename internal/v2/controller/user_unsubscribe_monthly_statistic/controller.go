package user_unsubscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_unsubscribe_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_unsubscribe_monthly_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user_unsubscribe_monthly_statistic.Resolver
}

func New(resolver user_unsubscribe_monthly_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSUserUnsubscribeMonthlyStatistic 獲取用戶退訂統計月報表
// @Summary 獲取用戶退訂統計月報表
// @Description 獲取用戶退訂統計月報表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} user_unsubscribe_monthly_statistic.APIGetCMSUserUnsubscribeStatisticOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/user/unsubscribe [GET]
func (c controller) GetCMSUserUnsubscribeMonthlyStatistic(ctx *gin.Context) {
	var input model.APIGetCMSUserUnsubscribeStatisticInput
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSUserSubscribeStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}
