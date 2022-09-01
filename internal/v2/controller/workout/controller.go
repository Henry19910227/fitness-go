package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
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

// CreateUserWorkout 創建個人訓練
// @Summary 創建個人訓練
// @Description 創建個人訓練
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

// DeleteUserWorkout 刪除個人訓練
// @Summary 刪除個人訓練
// @Description 刪除個人訓練
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

// UpdateUserWorkout 更新個人訓練
// @Summary 更新個人訓練
// @Description 前導音檔 : {Base URL}/v2/resource/workout/start_audio/{Filename} 結束音檔 : {Base URL}/v2/resource/workout/end_audio/{Filename}
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param name formData string false "訓練名稱"
// @Param equipment formData string false "所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"
// @Param start_audio formData file false "訓練前導音檔"
// @Param end_audio formData file false "訓練結束音檔"
// @Success 200 {object} workout.APIUpdateUserWorkoutOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id} [PATCH]
func (c *controller) UpdateUserWorkout(ctx *gin.Context) {
	input := model.APIUpdateUserWorkoutInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	//獲取訓練前導音檔
	file, fileHeader, _ := ctx.Request.FormFile("start_audio")
	if file != nil {
		input.Form.StartAudio = &fileModel.Input{}
		input.Form.StartAudio.Named = fileHeader.Filename
		input.Form.StartAudio.Data = file
	}
	//獲取訓練結束音檔
	file, fileHeader, _ = ctx.Request.FormFile("end_audio")
	if file != nil {
		input.Form.EndAudio = &fileModel.Input{}
		input.Form.EndAudio.Named = fileHeader.Filename
		input.Form.EndAudio.Data = file
	}
	output := c.resolver.APIUpdateUserWorkout(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}
