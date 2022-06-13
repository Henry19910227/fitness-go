package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type workoutset struct {
	Base
	workoutSetService service.WorkoutSet
}

func NewWorkoutSet(baseGroup *gin.RouterGroup,
	workoutSetService service.WorkoutSet,
	userMidd midd.User,
	courseMidd midd.Course)  {

	baseGroup.StaticFS("/resource/workout_set/start_audio", http.Dir("./volumes/storage/workout_set/start_audio"))
	baseGroup.StaticFS("/resource/workout_set/progress_audio", http.Dir("./volumes/storage/workout_set/progress_audio"))
	set := workoutset{workoutSetService: workoutSetService}

	baseGroup.POST("/workout_set/:workout_set_id/start_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.UploadWorkoutSetStartAudio)

	baseGroup.POST("/workout_set/:workout_set_id/progress_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.UploadWorkoutSetProgressAudio)

	baseGroup.POST("/workout_set/:workout_set_id/duplicate",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.DuplicateWorkoutSet)

	baseGroup.DELETE("/workout_set/:workout_set_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.DeleteWorkoutSet)

	baseGroup.PATCH("/workout_set/:workout_set_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.UpdateWorkoutSet)

	baseGroup.DELETE("/workout_set/:workout_set_id/start_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.DeleteWorkoutSetStartAudio)

	baseGroup.DELETE("/workout_set/:workout_set_id/progress_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		set.DeleteWorkoutSetProgressAudio)
}

// UpdateWorkoutSet 修改訓練組
// @Summary 修改訓練組
// @Description 修改訓練組
// @Tags WorkoutSet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param json_body body validator.UpdateWorkoutSetBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutSet} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /v1/workout_set/{workout_set_id} [PATCH]
func (w *workoutset) UpdateWorkoutSet(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	var body validator.UpdateWorkoutSetBody
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	set, err := w.workoutSetService.UpdateWorkoutSet(c, uri.WorkoutSetID, &dto.UpdateWorkoutSetParam{
		AutoNext: body.AutoNext,
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
// @Tags WorkoutSet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutSetID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/workout_set/{workout_set_id} [DELETE]
func (w *workoutset) DeleteWorkoutSet(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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
// @Description 下載訓練組前導語音 : https://www.fitopia-hub.tk/api/v1/resource/workout_set/start_audio/{語音檔案名}
// @Tags WorkoutSet_v1
// @Security fitness_token
// @Accept mpfd
// @Param workout_set_id path int64 true "訓練組id"
// @Param start_audio formData file true "前導語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutAudio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/workout_set/{workout_set_id}/start_audio [POST]
func (w *workoutset) UploadWorkoutSetStartAudio(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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

// DeleteWorkoutSetStartAudio 刪除訓練組前導語音
// @Summary 刪除訓練組前導語音
// @Description 刪除訓練組前導語音
// @Tags WorkoutSet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/workout_set/{workout_set_id}/start_audio [DELETE]
func (w *workoutset) DeleteWorkoutSetStartAudio(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.workoutSetService.DeleteWorkoutSetStartAudio(c, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, nil, "success")
}

// UploadWorkoutSetProgressAudio 上傳訓練組進行中語音
// @Summary 上傳訓練組進行中語音
// @Description 下載訓練組進行中語音 : https://www.fitopia-hub.tk/api/v1/resource/workout_set/progress_audio/{語音檔案名}
// @Tags WorkoutSet_v1
// @Security fitness_token
// @Accept mpfd
// @Param workout_set_id path int64 true "訓練組id"
// @Param progress_audio formData file true "進行中語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutAudio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/workout_set/{workout_set_id}/progress_audio [POST]
func (w *workoutset) UploadWorkoutSetProgressAudio(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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

// DeleteWorkoutSetProgressAudio 刪除訓練組進行中語音
// @Summary 刪除訓練組進行中語音
// @Description 刪除訓練組進行中語音
// @Tags WorkoutSet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/workout_set/{workout_set_id}/progress_audio [DELETE]
func (w *workoutset) DeleteWorkoutSetProgressAudio(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.workoutSetService.DeleteWorkoutSetProgressAudio(c, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, nil, "success")
}

// DuplicateWorkoutSet 複製訓練組
// @Summary 複製訓練組
// @Description 複製訓練組
// @Tags WorkoutSet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param json_body body validator.DuplicateWorkoutSetBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "複製成功!"
// @Failure 400 {object} model.ErrorResult "複製失敗"
// @Router /v1/workout_set/{workout_set_id}/duplicate [POST]
func (w *workoutset) DuplicateWorkoutSet(c *gin.Context) {
	var uri validator.WorkoutSetIDUri
	var body validator.DuplicateWorkoutSetBody
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	sets, err := w.workoutSetService.DuplicateWorkoutSets(c, uri.WorkoutSetID, body.DuplicateCount)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, sets, "duplicate success!")
}