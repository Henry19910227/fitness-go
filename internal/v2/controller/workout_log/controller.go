package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver workout_log.Resolver
}

func New(resolver workout_log.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateUserWorkoutLog 創建個人訓練紀錄
// @Summary 創建個人訓練紀錄
// @Description 創建個人訓練紀錄
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body workout_log.APICreateUserWorkoutLogBody true "輸入參數"
// @Success 200 {object} workout_log.APICreateUserWorkoutLogOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id}/workout_log [POST]
func (c *controller) CreateUserWorkoutLog(ctx *gin.Context) {
	var input model.APICreateUserWorkoutLogInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserWorkoutLog(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
