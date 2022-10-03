package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver workout_set_order.Resolver
}

func New(resolver workout_set_order.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdateUserWorkoutSetOrders 更新用戶個人訓練組排序
// @Summary 更新用戶個人訓練組排序
// @Description 更新用戶個人訓練組排序
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body workout_set_order.APIUpdateUserWorkoutSetOrderBody true "輸入參數"
// @Success 200 {object} workout_set_order.APIUpdateUserWorkoutSetOrdersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id}/workout_set_orders [PUT]
func (c *controller) UpdateUserWorkoutSetOrders(ctx *gin.Context) {
	var input model.APIUpdateUserWorkoutSetOrdersInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateUserWorkoutSetOrders(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateTrainerWorkoutSetOrders 更新教練訓練組排序
// @Summary 更新教練訓練組排序
// @Description 更新教練訓練組排序
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body workout_set_order.APIUpdateTrainerWorkoutSetOrderBody true "輸入參數"
// @Success 200 {object} workout_set_order.APIUpdateTrainerWorkoutSetOrdersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/workout/{workout_id}/workout_set_orders [PUT]
func (c *controller) UpdateTrainerWorkoutSetOrders(ctx *gin.Context) {
	var input model.APIUpdateTrainerWorkoutSetOrdersInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateTrainerWorkoutSetOrders(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
