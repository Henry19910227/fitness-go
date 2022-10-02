package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
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
		optional.WorkoutIDField
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

// CreateUserWorkoutSetByDuplicate 複製個人訓練組
// @Summary 複製個人訓練組
// @Description 複製個人訓練組
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param json_body body workout_set.APICreateUserWorkoutSetByDuplicateBody true "輸入參數"
// @Success 200 {object} workout_set.APICreateUserWorkoutSetByDuplicateOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout_set/{workout_set_id}/duplicate [POST]
func (c *controller) CreateUserWorkoutSetByDuplicate(ctx *gin.Context) {
	var input model.APICreateUserWorkoutSetByDuplicateInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserWorkoutSetByDuplicate(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// CreateUserRestSet 創建個人休息組
// @Summary 創建個人休息組
// @Description 創建個人休息組
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} workout_set.APICreateUserRestSetOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout/{workout_id}/rest_set [POST]
func (c *controller) CreateUserRestSet(ctx *gin.Context) {
	var input model.APICreateUserRestSetInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateUserRestSet(&input)
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

// UpdateUserWorkoutSet 更新個人訓練組
// @Summary 更新個人訓練組
// @Description 前導音檔 : {Base URL}/v2/resource/workout_set/start_audio/{Filename}.mp3 進行中音檔 : {Base URL}/v2/resource/workout_set/progress_audio/{Filename}.mp3
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param auto_next formData string false "自動下一組(Y:是/N:否)"
// @Param remark formData string false "備註(1~40)"
// @Param weight formData float64 false "重量(公斤 0.01~999.99)"
// @Param reps formData int false "次數(1~999)"
// @Param distance formData float64 false "距離(公里 0.01~999.99)"
// @Param duration formData int false "時長(秒 1~38439)"
// @Param incline formData float64 false "坡度(0.01~999.99)"
// @Param start_audio formData file false "訓練前導音檔"
// @Param progress_audio formData file false "訓練中音檔"
// @Success 200 {object} workout_set.APIUpdateUserWorkoutSetOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout_set/{workout_set_id} [PATCH]
func (c *controller) UpdateUserWorkoutSet(ctx *gin.Context) {
	input := model.APIUpdateUserWorkoutSetInput{}
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
	//獲取訓練中音檔
	file, fileHeader, _ = ctx.Request.FormFile("progress_audio")
	if file != nil {
		input.Form.ProgressAudio = &fileModel.Input{}
		input.Form.ProgressAudio.Named = fileHeader.Filename
		input.Form.ProgressAudio.Data = file
	}
	output := c.resolver.APIUpdateUserWorkoutSet(ctx.MustGet("tx").(*gorm.DB), &input)
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

// DeleteUserWorkoutSetStartAudio 刪除個人訓練組前導語音
// @Summary 刪除個人訓練組前導語音
// @Description 刪除個人訓練組前導語音
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} workout_set.APIDeleteUserWorkoutSetStartAudioOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout_set/{workout_set_id}/start_audio [DELETE]
func (c *controller) DeleteUserWorkoutSetStartAudio(ctx *gin.Context) {
	var input model.APIDeleteUserWorkoutSetStartAudioInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserWorkoutSetStartAudio(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteUserWorkoutSetProgressAudio 刪除個人訓練組進行中語音
// @Summary 刪除個人訓練組進行中語音
// @Description 刪除個人訓練組進行中語音
// @Tags 用戶個人課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} workout_set.APIDeleteUserWorkoutSetProgressAudioOutput "0:Success/ 9000:Bad Request/ 9005:Invalid Token/ 9006:Permission denied"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/workout_set/{workout_set_id}/progress_audio [DELETE]
func (c *controller) DeleteUserWorkoutSetProgressAudio(ctx *gin.Context) {
	var input model.APIDeleteUserWorkoutSetProgressAudioInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteUserWorkoutSetProgressAudio(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerWorkoutSets 獲取教練訓練組列表
// @Summary 獲取教練訓練組列表
// @Description 獲取教練訓練組列表
// @Tags 教練課表_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} workout_set.APIGetTrainerWorkoutSetsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/workout/{workout_id}/workout_sets [GET]
func (c *controller) GetTrainerWorkoutSets(ctx *gin.Context) {
	var input model.APIGetTrainerWorkoutSetsInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetTrainerWorkoutSets(&input)
	ctx.JSON(http.StatusOK, output)
}
