package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type WorkoutLog struct {
	Base
	workoutService    service.Workout
	workoutSetService service.WorkoutSet
	workoutLogService service.WorkoutLog
}

func NewWorkoutLog(baseGroup *gin.RouterGroup, workoutLogService service.WorkoutLog, userMidd midd.User) {
	workout := WorkoutLog{
		workoutLogService: workoutLogService,
	}
	baseGroup.GET("/workout_logs",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workout.GetWorkoutLogSummaries)
	baseGroup.GET("/workout_log/:workout_log_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workout.GetWorkoutLog)
	baseGroup.DELETE("/workout_log/:workout_log_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workout.DeleteWorkoutLog)
}

// GetWorkoutLog 獲取訓練紀錄詳細
// @Summary 獲取訓練紀錄詳細
// @Description 獲取訓練紀錄詳細
// @Tags History
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_log_id path int64 true "訓練記錄id"
// @Success 200 {object} model.SuccessResult{data=dto.WorkoutLog} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /workout_log/{workout_log_id} [GET]
func (p *WorkoutLog) GetWorkoutLog(c *gin.Context) {
	var uri validator.WorkoutLogIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workoutLog, e := p.workoutLogService.GetWorkoutLog(c, uri.WorkoutLogID)
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, workoutLog, "success!")
}

// GetWorkoutLogSummaries 以日期區間獲取訓練記錄
// @Summary 以日期區間獲取訓練記錄
// @Description 以日期區間獲取訓練記錄，用於獲取歷史首頁資料
// @Tags History
// @Accept json
// @Produce json
// @Security fitness_token
// @Param start_date query string true "區間開始日期 YYYY-MM-DD"
// @Param end_date query string true "區間結束日期 YYYY-MM-DD"
// @Success 200 {object} model.SuccessResult{date=[]dto.WorkoutLogSummary} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout_logs [GET]
func (p *WorkoutLog) GetWorkoutLogSummaries(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var query validator.GetWorkoutLogSummariesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workoutLogs, e := p.workoutLogService.GetWorkoutLogSummaries(c, uid, query.StartDate, query.EndDate)
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, workoutLogs, "success!")
}

// DeleteWorkoutLog 刪除訓練紀錄
// @Summary 刪除訓練紀錄
// @Description 用於歷史頁面刪除訓練紀錄
// @Tags History
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_log_id path int64 true "訓練記錄id"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /workout_log/{workout_log_id} [DELETE]
func (p *WorkoutLog) DeleteWorkoutLog(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var uri validator.WorkoutLogIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	e := p.workoutLogService.DeleteWorkoutLog(c, uid, uri.WorkoutLogID)
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, nil, "success!")
}