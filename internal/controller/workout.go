package controller

import (
	"github.com/Henry19910227/fitness-go/internal/access"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type workout struct {
	Base
	workoutService    service.Workout
	workoutSetService service.WorkoutSet
	courseAccess      access.Course
}

func NewWorkout(baseGroup *gin.RouterGroup,
	workoutService service.Workout,
	workoutSetService service.WorkoutSet,
	courseAccess access.Course,
	userMiddleware gin.HandlerFunc) {
	baseGroup.StaticFS("/resource/workout/audio", http.Dir("./volumes/storage/workout/audio"))
	workout := workout{workoutService: workoutService, workoutSetService: workoutSetService, courseAccess: courseAccess}
	planGroup := baseGroup.Group("/workout")
	planGroup.Use(userMiddleware)
	planGroup.PATCH("/:workout_id", workout.UpdateWorkout)
	planGroup.DELETE("/:workout_id", workout.DeleteWorkout)
	planGroup.POST("/:workout_id/start_audio", workout.UploadWorkoutStartAudio)
	planGroup.POST("/:workout_id/end_audio", workout.UploadWorkoutEndAudio)
	planGroup.POST("/:workout_id/rest_set", workout.CreateRestSet)
}

// UpdateWorkout 修改訓練
// @Summary 修改訓練
// @Description 修改訓練
// @Tags Workout
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.UpdateWorkoutBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=workoutdto.Workout} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /workout/{workout_id} [PATCH]
func (w *workout) UpdateWorkout(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	var body validator.UpdateWorkoutBody

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
	if err := w.courseAccess.CheckEditAllowByWorkoutID(c, header.Token, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
    workout, err := w.workoutService.UpdateWorkout(c, uri.WorkoutID, &workoutdto.UpdateWorkoutParam{
		Name: body.Name,
		Equipment: body.Equipment,
	})
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, workout, "update success!")
}

// DeleteWorkout 刪除訓練
// @Summary 刪除訓練
// @Description 刪除訓練
// @Tags Workout
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutID} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout/{workout_id} [DELETE]
func (w *workout) DeleteWorkout(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.courseAccess.CheckEditAllowByWorkoutID(c, header.Token, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	data, err := w.workoutService.DeleteWorkout(c, uri.WorkoutID)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, data, "delete success!")
}

// UploadWorkoutStartAudio 上傳訓練前導語音
// @Summary 上傳訓練前導語音
// @Description 下載前導語音 : https://www.fitness-app.tk/api/v1/resource/workout/audio/{語音檔案名}
// @Tags Workout
// @Security fitness_user_token
// @Accept mpfd
// @Param workout_id path int64 true "訓練id"
// @Param start_audio formData file true "前導語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=workoutdto.Audio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /workout/{workout_id}/start_audio [POST]
func (w *workout) UploadWorkoutStartAudio(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.courseAccess.CheckEditAllowByWorkoutID(c, header.Token, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	file, fileHeader, err := c.Request.FormFile("start_audio")
	if err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := w.workoutService.UploadWorkoutStartAudio(c, uri.WorkoutID, fileHeader.Filename, file)
	if e != nil {
		w.JSONErrorResponse(c, e)
		return
	}
	w.JSONSuccessResponse(c, result, "upload success")
}

// UploadWorkoutEndAudio 上傳訓練結束語音
// @Summary 上傳訓練結束語音
// @Description 上傳訓練結束語音 : https://www.fitness-app.tk/api/v1/resource/workout/audio/{語音檔案名}
// @Tags Workout
// @Security fitness_user_token
// @Accept mpfd
// @Param workout_id path int64 true "訓練id"
// @Param end_audio formData file true "結束語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=workoutdto.Audio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /workout/{workout_id}/end_audio [POST]
func (w *workout) UploadWorkoutEndAudio(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.courseAccess.CheckEditAllowByWorkoutID(c, header.Token, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	file, fileHeader, err := c.Request.FormFile("end_audio")
	if err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := w.workoutService.UploadWorkoutEndAudio(c, uri.WorkoutID, fileHeader.Filename, file)
	if e != nil {
		w.JSONErrorResponse(c, e)
		return
	}
	w.JSONSuccessResponse(c, result, "upload success")
}

// CreateRestSet 新增休息組
// @Summary 新增休息組
// @Description 新增休息組
// @Tags Workout
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutSet} "新增成功!"
// @Failure 400 {object} model.ErrorResult "新增失敗"
// @Router /workout/{workout_id}/rest_set [POST]
func (w *workout) CreateRestSet(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.courseAccess.CheckEditAllowByWorkoutID(c, header.Token, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	set, err := w.workoutSetService.CreateRestSet(c, uri.WorkoutID)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, set, "create success!")
}