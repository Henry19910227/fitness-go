package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/receipt/api_get_cms_order_receipts"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/receipt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver receipt.Resolver
}

func New(resolver receipt.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSOrderReceipts 獲取訂單收據列表
// @Summary 獲取訂單收據列表
// @Description 獲取訂單收據列表
// @Tags CMS訂單管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_id path string true "訂單ID"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} api_get_cms_order_receipts.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/order/{order_id}/receipts [GET]
func (c *controller) GetCMSOrderReceipts(ctx *gin.Context) {
	input := api_get_cms_order_receipts.Input{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSOrderReceipts(&input)
	ctx.JSON(http.StatusOK, output)
}
