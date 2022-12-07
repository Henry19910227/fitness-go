package rda

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_update_rda"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/rda"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver rda.Resolver
}

func New(resolver rda.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdateRDA 更新建議飲食攝取量
// @Summary 更新建議飲食攝取量
// @Description 更新建議飲食攝取量
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_update_rda.Body true "輸入參數"
// @Success 200 {object} api_update_rda.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/rda [PUT]
func (c *controller) UpdateRDA(ctx *gin.Context) {
	input := api_update_rda.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateRDA(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
