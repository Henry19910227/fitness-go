package favorite_action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_action"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver favorite_action.Resolver
}

func New(resolver favorite_action.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateFavoriteAction 收藏動作
// @Summary 收藏動作
// @Description 收藏動作
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} favorite_action.APICreateFavoriteActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/action/{action_id} [POST]
func (c *controller) CreateFavoriteAction(ctx *gin.Context) {
	input := model.APICreateFavoriteActionInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateFavoriteAction(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteFavoriteAction 刪除動作收藏
// @Summary 刪除動作收藏
// @Description 刪除動作收藏
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} favorite_action.APIDeleteFavoriteActionOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/action/{action_id} [DELETE]
func (c *controller) DeleteFavoriteAction(ctx *gin.Context) {
	input := model.APIDeleteFavoriteActionInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteFavoriteAction(&input)
	ctx.JSON(http.StatusOK, output)
}
