package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
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

// GetCMSReceipts 獲取收據列表
// @Summary 獲取收據列表
// @Description 獲取收據列表
// @Tags CMS訂單管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} receipt.APIGetCMSReceiptsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/receipts [GET]
func (c controller) GetCMSReceipts(ctx *gin.Context) {
	input := model.APIGetCMSReceiptsInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSReceipts(&input)
	ctx.JSON(http.StatusOK, output)
}
