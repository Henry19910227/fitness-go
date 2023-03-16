package user_subscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic/api_get_cms_user_subscribe_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_subscribe_monthly_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user_subscribe_monthly_statistic.Resolver
}

func New(resolver user_subscribe_monthly_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSUserSubscribeMonthlyStatistic 獲取用戶訂閱統計月報表
// @Summary 獲取用戶訂閱統計月報表
// @Description 獲取用戶訂閱統計月報表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} user_subscribe_monthly_statistic.APIGetCMSUserSubscribeStatisticOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/user/subscribe [GET]
func (c *controller) GetCMSUserSubscribeMonthlyStatistic(ctx *gin.Context) {
	var input api_get_cms_user_subscribe_statistic.Input
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSUserSubscribeStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}

func (c *controller) Statistic() {
	c.resolver.Statistic()
}
