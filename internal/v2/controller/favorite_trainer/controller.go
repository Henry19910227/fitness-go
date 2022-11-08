package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer/api_create_favorite_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_trainer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver favorite_trainer.Resolver
}

func New(resolver favorite_trainer.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateFavoriteTrainer 收藏教練
// @Summary 收藏教練
// @Description 收藏教練
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} api_create_favorite_trainer.Output "Success"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/trainer/{user_id} [POST]
func (c *controller) CreateFavoriteTrainer(ctx *gin.Context) {
	input := api_create_favorite_trainer.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateFavoriteTrainer(&input)
	ctx.JSON(http.StatusOK, output)
}
