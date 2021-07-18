package controller

import (
	"github.com/Henry19910227/fitness-go/internal/access"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type workoutset struct {
	Base
	workoutSetService service.WorkoutSet
	workoutSetAccess  access.WorkoutSet
	trainerAccess access.Trainer
}

func NewWorkoutSet(baseGroup *gin.RouterGroup,
	workoutSetService service.WorkoutSet,
	workoutSetAccess access.WorkoutSet,
	trainerAccess access.Trainer,
	userMiddleware gin.HandlerFunc)  {

	baseGroup.StaticFS("/resource/workout_set/audio", http.Dir("./volumes/storage/workout_set/audio"))
	baseGroup.StaticFS("/resource/workout_set/progress_audio", http.Dir("./volumes/storage/workout_set/progress_audio"))
	set := workoutset{workoutSetService: workoutSetService,
		workoutSetAccess: workoutSetAccess,
		trainerAccess: trainerAccess}
	setGroup := baseGroup.Group("/workout_set")
	setGroup.Use(userMiddleware)
	setGroup.PATCH("/:workout_set_id", set.UpdateWorkoutSet)
	setGroup.DELETE("/:workout_set_id", set.DeleteWorkoutSet)
	setGroup.POST("/:workout_set_id/start_audio", set.UploadWorkoutSetStartAudio)
	setGroup.POST("/:workout_set_id/progress_audio", set.UploadWorkoutSetProgressAudio)
}

// UpdateWorkoutSet 修改訓練組
// @Summary 修改訓練組
// @Description 修改訓練組
// @Tags WorkoutSet
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param json_body body validator.UpdateWorkoutSetBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutSet} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /workout_set/{workout_set_id} [PATCH]
func (w *workoutset) UpdateWorkoutSet(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutSetIDUri
	var body validator.UpdateWorkoutSetBody
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.trainerAccess.StatusVerify(c, header.Token); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	if err := w.workoutSetAccess.UpdateVerifyByWorkoutSetID(c, header.Token, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	set, err := w.workoutSetService.UpdateWorkoutSet(c, uri.WorkoutSetID, &workoutdto.UpdateWorkoutSetParam{
		AutoNext: body.AutoNext,
		StartAudio: body.StartAudio,
		Remark: body.Remark,
		Weight: body.Weight,
		Reps: body.Reps,
		Distance: body.Distance,
		Duration: body.Duration,
		Incline: body.Incline,
	})
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, set, "update success!")
}

// DeleteWorkoutSet 刪除訓練組
// @Summary 刪除訓練組
// @Description 刪除訓練組
// @Tags WorkoutSet
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutSetID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /workout_set/{workout_set_id} [DELETE]
func (w *workoutset) DeleteWorkoutSet(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.trainerAccess.StatusVerify(c, header.Token); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	if err := w.workoutSetAccess.UpdateVerifyByWorkoutSetID(c, header.Token, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	result, err := w.workoutSetService.DeleteWorkoutSet(c, uri.WorkoutSetID)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, result, "delete success!")
}

// UploadWorkoutSetStartAudio 上傳訓練組前導語音
// @Summary 上傳訓練組前導語音
// @Description 下載訓練組前導語音 : https://www.fitness-app.tk/api/v1/resource/workout_set/audio/{語音檔案名}
// @Tags WorkoutSet
// @Security fitness_user_token
// @Accept mpfd
// @Param workout_set_id path int64 true "訓練組id"
// @Param start_audio formData file true "前導語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=workoutdto.Audio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /workout_set/{workout_set_id}/start_audio [POST]
func (w *workoutset) UploadWorkoutSetStartAudio(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.trainerAccess.StatusVerify(c, header.Token); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	if err := w.workoutSetAccess.UpdateVerifyByWorkoutSetID(c, header.Token, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	file, fileHeader, err := c.Request.FormFile("start_audio")
	if err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := w.workoutSetService.UploadWorkoutSetStartAudio(c, uri.WorkoutSetID, fileHeader.Filename, file)
	if e != nil {
		w.JSONErrorResponse(c, e)
		return
	}
	w.JSONSuccessResponse(c, result, "upload success")
}

// UploadWorkoutSetProgressAudio 上傳訓練組進行中語音
// @Summary 上傳訓練組進行中語音
// @Description 下載訓練組進行中語音 : https://www.fitness-app.tk/api/v1/resource/workout_set/progress_audio/{語音檔案名}
// @Tags WorkoutSet
// @Security fitness_user_token
// @Accept mpfd
// @Param workout_set_id path int64 true "訓練組id"
// @Param progress_audio formData file true "進行中語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=workoutdto.Audio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /workout_set/{workout_set_id}/progress_audio [POST]
func (w *workoutset) UploadWorkoutSetProgressAudio(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.trainerAccess.StatusVerify(c, header.Token); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	if err := w.workoutSetAccess.UpdateVerifyByWorkoutSetID(c, header.Token, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	file, fileHeader, err := c.Request.FormFile("progress_audio")
	if err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := w.workoutSetService.UploadWorkoutSetProgressAudio(c, uri.WorkoutSetID, fileHeader.Filename, file)
	if e != nil {
		w.JSONErrorResponse(c, e)
		return
	}
	w.JSONSuccessResponse(c, result, "upload success")
}