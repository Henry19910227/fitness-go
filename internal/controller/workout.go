package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type workout struct {
	Base
	workoutService    service.Workout
	workoutSetService service.WorkoutSet
}

func NewWorkout(baseGroup *gin.RouterGroup,
	workoutService service.Workout,
	workoutSetService service.WorkoutSet,
	userMidd midd.User,
	courseMidd midd.Course) {
	baseGroup.StaticFS("/resource/workout/start_audio", http.Dir("./volumes/storage/workout/start_audio"))
	baseGroup.StaticFS("/resource/workout/end_audio", http.Dir("./volumes/storage/workout/end_audio"))
	workout := workout{workoutService: workoutService,
		workoutSetService: workoutSetService}

	baseGroup.PATCH("/workout/:workout_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.UpdateWorkout)

	baseGroup.DELETE("/workout/:workout_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.DeleteWorkout)

	baseGroup.POST("/workout/:workout_id/start_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.UploadWorkoutStartAudio)

	baseGroup.POST("/workout/:workout_id/end_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.UploadWorkoutEndAudio)

	baseGroup.DELETE("/workout/:workout_id/start_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.DeleteWorkoutStartAudio)

	baseGroup.DELETE("/workout/:workout_id/end_audio",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.DeleteWorkoutEndAudio)

	baseGroup.POST("/workout/:workout_id/workout_set",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.CreateWorkoutSets)

	baseGroup.POST("/workout/:workout_id/rest_set",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		workout.CreateRestSet)

	baseGroup.GET("/workout/:workout_id/workout_sets",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		workout.GetWorkoutSets)

	baseGroup.PUT("/workout/:workout_id/order",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		workout.UpdateWorkoutSetOrders)
}

// UpdateWorkout 修改訓練
// @Summary 修改訓練
// @Description 修改訓練
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.UpdateWorkoutBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Workout} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /v1/workout/{workout_id} [PATCH]
func (w *workout) UpdateWorkout(c *gin.Context) {
	var uri validator.WorkoutIDUri
	var body validator.UpdateWorkoutBody

	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workout, err := w.workoutService.UpdateWorkout(c, uri.WorkoutID, &dto.UpdateWorkoutParam{
		Name:      body.Name,
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
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutID} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/workout/{workout_id} [DELETE]
func (w *workout) DeleteWorkout(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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
// @Description 下載前導語音 : https://www.fitopia-hub.tk/api/v1/resource/workout/start_audio/{語音檔案名}
// @Tags Workout_v1
// @Security fitness_token
// @Accept mpfd
// @Param workout_id path int64 true "訓練id"
// @Param start_audio formData file true "前導語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutAudio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/workout/{workout_id}/start_audio [POST]
func (w *workout) UploadWorkoutStartAudio(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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
// @Description 上傳訓練結束語音 : https://www.fitopia-hub.tk/api/v1/resource/workout/end_audio/{語音檔案名}
// @Tags Workout_v1
// @Security fitness_token
// @Accept mpfd
// @Param workout_id path int64 true "訓練id"
// @Param end_audio formData file true "結束語音"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutAudio} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/workout/{workout_id}/end_audio [POST]
func (w *workout) UploadWorkoutEndAudio(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
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

// DeleteWorkoutStartAudio 刪除訓練前導語音
// @Summary 刪除訓練前導語音
// @Description 刪除訓練前導語音
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/workout/{workout_id}/start_audio [DELETE]
func (w *workout) DeleteWorkoutStartAudio(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.workoutService.DeleteWorkoutStartAudio(c, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, nil, "success")
}

// DeleteWorkoutEndAudio 刪除訓練結束語音
// @Summary 刪除訓練結束語音
// @Description 刪除訓練結束語音
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/workout/{workout_id}/end_audio [DELETE]
func (w *workout) DeleteWorkoutEndAudio(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := w.workoutService.DeleteWorkoutEndAudio(c, uri.WorkoutID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, nil, "success")
}

// CreateWorkoutSets 新增訓練組
// @Summary 新增訓練組
// @Description 新增訓練組
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.CreateWorkoutSetBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "新增成功!"
// @Failure 400 {object} model.ErrorResult "新增失敗"
// @Router /v1/workout/{workout_id}/workout_set [POST]
func (w *workout) CreateWorkoutSets(c *gin.Context) {
	var uri validator.WorkoutIDUri
	var body validator.CreateWorkoutSetBody
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	sets, err := w.workoutSetService.CreateWorkoutSets(c, uri.WorkoutID, body.ActionIDs)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, sets, "create success!")
}

// CreateRestSet 新增休息組
// @Summary 新增休息組
// @Description 新增休息組
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutSet} "新增成功!"
// @Failure 400 {object} model.ErrorResult "新增失敗"
// @Router /v1/workout/{workout_id}/rest_set [POST]
func (w *workout) CreateRestSet(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	set, err := w.workoutSetService.CreateRestSet(c, uri.WorkoutID)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, set, "create success!")
}

// GetWorkoutSets 取得訓練內的訓練組列表
// @Summary  取得訓練內的訓練組列表
// @Description  取得訓練內的訓練組列表
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/workout/{workout_id}/workout_sets [GET]
func (w *workout) GetWorkoutSets(c *gin.Context) {
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
	sets, err := w.workoutSetService.GetWorkoutSets(c, uri.WorkoutID, nil)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, sets, "success!")
}

// UpdateWorkoutSetOrders 修改訓練組的順序
// @Summary 修改訓練組的順序
// @Description 修改訓練組的順序
// @Tags Workout_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.UpdateWorkoutSetOrderBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /v1/workout/{workout_id}/order [PUT]
func (w *workout) UpdateWorkoutSetOrders(c *gin.Context) {
	var uri validator.WorkoutIDUri
	var body validator.UpdateWorkoutSetOrderBody
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var orders []*dto.WorkoutSetOrder
	for _, data := range body.Orders {
		order := dto.WorkoutSetOrder{
			WorkoutSetID: data.WorkoutSetID,
			Seq:          data.Seq,
		}
		orders = append(orders, &order)
	}
	if err := w.workoutSetService.UpdateWorkoutSetOrders(c, uri.WorkoutID, orders); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, nil, "update success!")
}
