package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type PlanAsset struct {
	Base
	planService    service.Plan
	workoutService service.Workout
}

func NewPlanAsset(baseGroup *gin.RouterGroup, planService service.Plan, workoutService service.Workout, planMidd midd.Plan, userMidd midd.User) {
	plan := PlanAsset{
		planService:    planService,
		workoutService: workoutService,
	}
	baseGroup.GET("/plan_asset/:plan_id/workouts",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		planMidd.CourseStatusVerify(planService.GetPlanStatus, []global.CourseStatus{global.Sale}),
		plan.GetWorkouts)
}

// GetWorkouts 獲取訓練列表
// @Summary 獲取訓練列表
// @Description 獲取訓練列表
// @Tags Exercise_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutAsset} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/plan_asset/{plan_id}/workouts [GET]
func (p *PlanAsset) GetWorkouts(c *gin.Context) {
	uid, e := p.GetUID(c)
	if e != nil {
		p.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := p.workoutService.GetWorkoutAssets(c, uid, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, workouts, "success!")
}
