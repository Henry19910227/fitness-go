package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver workoutSet.Resolver
}

func New(resolver workoutSet.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSWorkoutSets 獲取訓練組列表
// @Summary 獲取訓練組列表
// @Description 獲取訓練組列表
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練ID"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} workout_set.APIGetCMSWorkoutSetsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/workout/{workout_id}/workout_sets [GET]
func (c *controller) GetCMSWorkoutSets(ctx *gin.Context) {
	var uri struct {
		model.WorkoutIDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	type pagingInput paging.Input
	var query struct {
		pagingInput
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSWorkoutSetsInput{}
	input.WorkoutID = uri.WorkoutID
	input.Page = query.Page
	input.Size = query.Size
	output := c.resolver.APIGetCMSWorkoutSets(&input)
	ctx.JSON(http.StatusOK, output)
}
