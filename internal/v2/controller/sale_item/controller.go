package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/sale_item/api_get_sale_items"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/sale_item"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver sale_item.Resolver
}

func New(resolver sale_item.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetSaleItems 取得銷售項目清單
// @Summary 取得銷售項目清單
// @Description 取得銷售項目清單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_get_sale_items.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/sale_items [GET]
func (c *controller) GetSaleItems(ctx *gin.Context) {
	input := api_get_sale_items.Input{}
	output := c.resolver.APIGetSaleItems(&input)
	ctx.JSON(http.StatusOK, output)
}
