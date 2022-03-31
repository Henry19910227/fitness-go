package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
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
// @Summary 獲取訓練組列表
// @Description 獲取訓練組列表
// @Tags Exercise
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout_asset/{workout_id}/workout_sets [GET]
func (p *WorkoutAsset) GetWorkoutSets(c *gin.Context) {
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := p.workoutSetService.GetWorkoutSets(c, uri.WorkoutID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, workouts, "success!")
}

// CreateWorkoutLog 創建訓練記錄
// @Summary 創建訓練記錄
// @Description 創建訓練記錄
// @Tags Exercise
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.CreateWorkoutLogBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSetLogTag} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout_asset/{workout_id}/workout_log [POST]
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
