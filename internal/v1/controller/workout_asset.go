package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type WorkoutAsset struct {
	Base
	workoutService    service.Workout
	workoutSetService service.WorkoutSet
	workoutLogService service.WorkoutLog
}

func NewWorkoutAsset(baseGroup *gin.RouterGroup, workoutService service.Workout, workoutSetService service.WorkoutSet, workoutLogService service.WorkoutLog, workoutMidd midd.Workout, userMidd midd.User) {
	workout := WorkoutAsset{
		workoutService:    workoutService,
		workoutSetService: workoutSetService,
		workoutLogService: workoutLogService,
	}
	baseGroup.GET("/workout_asset/:workout_id/workout_sets",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workoutMidd.CourseStatusVerify(workoutService.GetWorkoutStatus, []global.CourseStatus{global.Sale}),
		workout.GetWorkoutSets)

	baseGroup.POST("/workout_asset/:workout_id/workout_log",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workoutMidd.CourseStatusVerify(workoutService.GetWorkoutStatus, []global.CourseStatus{global.Sale}),
		workout.CreateWorkoutLog)
}

// GetWorkoutSets 獲取訓練組列表
// @Summary 獲取訓練組列表 (API已經過時，更新為 /v2/user/workout/{workout_id}/workout_sets [GET])
// @Description 獲取訓練組列表
// @Tags Exercise_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/workout_asset/{workout_id}/workout_sets [GET]
func (p *WorkoutAsset) GetWorkoutSets(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, e := p.workoutSetService.GetWorkoutSets(c, uri.WorkoutID, &uid)
	if err != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, workouts, "success!")
}

// CreateWorkoutLog 創建訓練記錄
// @Summary 創建訓練記錄 (API已經過時，更新為 /v2/user/workout/{workout_id}/workout_log [POST])
// @Description 創建訓練記錄
// @Tags Exercise_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.CreateWorkoutLogBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSetLogTag} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/workout_asset/{workout_id}/workout_log [POST]
func (p *WorkoutAsset) CreateWorkoutLog(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.CreateWorkoutLogBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workoutSetLogTags, e := p.workoutLogService.CreateWorkoutLog(c, uid, uri.WorkoutID, &dto.WorkoutLogParam{
		Duration:       body.Duration,
		Intensity:      body.Intensity,
		Place:          body.Place,
		WorkoutSetLogs: body.WorkoutSetLogs,
	})
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, workoutSetLogTags, "success!")
}
