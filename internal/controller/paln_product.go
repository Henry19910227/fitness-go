package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type PlanProduct struct {
	Base
	planService service.Plan
	workoutService service.Workout
}

func NewPlanProduct(baseGroup *gin.RouterGroup, planService service.Plan, workoutService service.Workout, planMidd midd.Plan, userMidd midd.User) {
	plan := PlanProduct{
		planService: planService,
		workoutService: workoutService,
	}
	baseGroup.GET("/plan_product/:plan_id/workouts",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		planMidd.CourseStatusVerify(planService.GetPlanStatus, []global.CourseStatus{global.Sale}),
		plan.GetWorkouts)
}

// GetWorkouts 獲取訓練列表
// @Summary 獲取訓練列表
// @Description 獲取訓練列表
// @Tags PlanProduct
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutProduct} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /plan_product/{plan_id}/workouts [GET]
func (p *PlanProduct) GetWorkouts(c *gin.Context) {
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := p.workoutService.GetWorkoutProductsByPlanID(c, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, workouts, "success!")
}