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

// CreateUserWorkout 創建個人課表訓練
// @Summary 創建個人課表訓練
// @Description 創建個人課表訓練
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body workout.APICreateUserWorkoutBody true "輸入參數"
// @Success 200 {object} workout.APICreateUserWorkoutOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/plan/{plan_id}/workout [POST]
func (c *controller) CreateUserWorkout(ctx *gin.Context) {
	var input model.APICreateUserWorkoutInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserWorkout(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserWorkout 刪除個人課表訓練
// @Summary 刪除個人課表訓練
// @Description 刪除個人課表訓練
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} workout.APIDeleteUserWorkoutOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id} [DELETE]
func (c *controller) DeleteUserWorkout(ctx *gin.Context) {
	var input model.APIDeleteUserWorkoutInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserWorkout(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserWorkouts 獲取用戶個人訓練列表
// @Summary 獲取用戶個人訓練列表
// @Description 獲取用戶個人訓練列表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} workout.APIGetUserWorkoutsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/plan/{plan_id}/workouts [GET]
func (c *controller) GetUserWorkouts(ctx *gin.Context) {
	input := model.APIGetUserWorkoutsInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserWorkouts(&input)
	ctx.JSON(http.StatusOK, output)
}
