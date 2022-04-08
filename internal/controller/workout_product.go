package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type WorkoutProduct struct {
	Base
	workoutService    service.Workout
	workoutSetService service.WorkoutSet
	workoutLogService service.WorkoutLog
}

func NewWorkoutProduct(baseGroup *gin.RouterGroup, workoutService service.Workout, workoutSetService service.WorkoutSet, workoutLogService service.WorkoutLog, workoutMidd midd.Workout, userMidd midd.User) {
	workout := WorkoutProduct{
		workoutService:    workoutService,
		workoutSetService: workoutSetService,
		workoutLogService: workoutLogService,
	}
	baseGroup.GET("/workout_product/:workout_id/workout_sets",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		workoutMidd.CourseStatusVerify(workoutService.GetWorkoutStatus, []global.CourseStatus{global.Sale}),
		workout.GetWorkoutSets)
}

// GetWorkoutSets 獲取訓練組列表(探索區課表)
// @Summary 獲取訓練組列表(探索區課表)
// @Description 獲取訓練組列表(探索區課表)
// @Tags Explore
// @Accept json
// @Produce json
// @Security fitness_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSet} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout_product/{workout_id}/workout_sets [GET]
func (p *WorkoutProduct) GetWorkoutSets(c *gin.Context) {
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
