package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver workout.Resolver
}

func New(resolver workout.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreatePersonalWorkout 創建個人課表訓練
// @Summary 創建個人課表訓練
// @Description 創建個人課表訓練
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body workout.APICreatePersonalWorkoutBody true "輸入參數"
// @Success 200 {object} workout.APICreatePersonalWorkoutOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/personal/plan/{plan_id}/workout [POST]
func (c *controller) CreatePersonalWorkout(ctx *gin.Context) {
	var input model.APICreatePersonalWorkoutInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreatePersonalWorkout(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
