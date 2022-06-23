package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/action"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver action.Resolver
}

func New(resolver action.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSActions 獲取動作列表
// @Summary 獲取動作列表
// @Description 獲取動作列表
// @Tags CMS內容管理_動作庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} action.APIGetCMSActionsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/actions [GET]
func (c *controller) GetCMSActions(ctx *gin.Context) {
	var query struct {
		paging.Input
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSActionsInput{}
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSActions(&input)
	ctx.JSON(http.StatusOK, output)
}
