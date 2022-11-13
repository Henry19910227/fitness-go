package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan/api_get_subscribe_plans"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/subscribe_plan"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver subscribe_plan.Resolver
}

func New(resolver subscribe_plan.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetSubscribePlans 取得訂閱方案清單
// @Summary 取得訂閱方案清單
// @Description 取得訂閱方案清單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_get_subscribe_plans.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/subscribe_plans [GET]
func (c *controller) GetSubscribePlans(ctx *gin.Context) {
	input := api_get_subscribe_plans.Input{}
	output := c.resolver.APIGetSubscribePlans(&input)
	ctx.JSON(http.StatusOK, output)
}
