package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver workoutSet.Resolver
}

func New(resolver workoutSet.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateUserWorkoutSets 創建個人訓練組
// @Summary 創建個人訓練組
// @Description 創建個人訓練組
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body workout_set.APICreateUserWorkoutSetsBody true "輸入參數"
// @Success 200 {object} workout_set.APICreateUserWorkoutSetsOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id}/workout_sets [POST]
func (c *controller) CreateUserWorkoutSets(ctx *gin.Context) {
	var input model.APICreateUserWorkoutSetsInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserWorkoutSets(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserWorkoutSet 刪除個人訓練組
// @Summary 刪除個人訓練組
// @Description 刪除個人訓練組
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} workout_set.APIDeleteUserWorkoutSetOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout_set/{workout_set_id} [DELETE]
func (c *controller) DeleteUserWorkoutSet(ctx *gin.Context) {
	var input model.APIDeleteUserWorkoutSetInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserWorkoutSet(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserWorkoutSets 獲取用戶個人訓練組列表
// @Summary 獲取用戶個人訓練組列表
// @Description 獲取用戶個人訓練組列表
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} workout_set.APIGetUserWorkoutSetsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id}/workout_sets [GET]
func (c *controller) GetUserWorkoutSets(ctx *gin.Context) {
	var input model.APIGetUserWorkoutSetsInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetUserWorkoutSets(&input)
	ctx.JSON(http.StatusOK, output)
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
