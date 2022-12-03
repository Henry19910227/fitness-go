package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner_order/api_update_cms_banner_orders"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/banner_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver banner_order.Resolver
}

func New(resolver banner_order.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdateCMSBannerOrders 更新banner排序
// @Summary 更新banner排序
// @Description 更新banner排序
// @Tags CMS內容管理_Banner_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_update_cms_banner_orders.Body true "輸入參數"
// @Success 200 {object} api_update_cms_banner_orders.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/banner_orders [PUT]
func (c *controller) UpdateCMSBannerOrders(ctx *gin.Context) {
	input := api_update_cms_banner_orders.Input{}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateCMSBannerOrders(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}