package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log/api_get_user_action_workout_set_logs"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set_log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver workout_set_log.Resolver
}

func New(resolver workout_set_log.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetUserActionWorkoutSetLogs 以日期獲取動作訓練組紀錄
// @Summary 以日期獲取動作訓練組紀錄
// @Description 以日期獲取動作訓練組紀錄
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作ID"
// @Param start_date query string true "區間開始日期 YYYY-MM-DD"
// @Param end_date query string true "區間結束日期 YYYY-MM-DD"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} api_get_user_action_workout_set_logs.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/action/{action_id}/workout_set_logs [GET]
func (c *controller) GetUserActionWorkoutSetLogs(ctx *gin.Context) {
	var input api_get_user_action_workout_set_logs.Input
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserActionWorkoutSetLogs(&input)
	ctx.JSON(http.StatusOK, output)
}
