package controller

import (
	"github.com/Henry19910227/fitness-go/internal/access"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Plan struct {
	Base
	planService    service.Plan
	workoutService service.Workout
	planAccess     access.Plan
	workoutAccess access.Workout
	trainerAccess  access.Trainer
}

func NewPlan(baseGroup *gin.RouterGroup,
	planService service.Plan,
	workoutService service.Workout,
	planAccess access.Plan,
	workoutAccess access.Workout,
	trainerAccess  access.Trainer,
	userMiddleware gin.HandlerFunc) {
	plan := Plan{planService: planService,
		workoutService: workoutService,
		planAccess: planAccess,
		workoutAccess: workoutAccess,
		trainerAccess: trainerAccess}
	planGroup := baseGroup.Group("/plan")
	planGroup.Use(userMiddleware)
	planGroup.PATCH("/:plan_id", plan.UpdatePlan)
	planGroup.DELETE("/:plan_id", plan.DeletePlan)
	planGroup.POST("/:plan_id/workout", plan.CreateWorkout)
	planGroup.GET("/:plan_id/workouts", plan.GetWorkouts)
}

// UpdatePlan 修改計畫
// @Summary 修改計畫
// @Description 修改計畫
// @Tags Plan
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body validator.UpdatePlanBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=plandto.Plan} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /plan/{plan_id} [PATCH]
func (p *Plan) UpdatePlan(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.PlanIDUri
	var body validator.UpdatePlanBody
	if err := c.ShouldBindHeader(&header); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.trainerAccess.StatusVerify(c, header.Token); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	if err := p.planAccess.UpdateVerifyByPlanID(c, header.Token, uri.PlanID); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	planData, err := p.planService.UpdatePlan(c, uri.PlanID, body.Name)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, planData, "update success")
}

// DeletePlan 刪除計畫
// @Summary 刪除計畫
// @Description 刪除計畫
// @Tags Plan
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=plandto.PlanID} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /plan/{plan_id} [DELETE]
func (p *Plan) DeletePlan(c *gin.Context)  {
	var header validator.TokenHeader
	var uri validator.PlanIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.trainerAccess.StatusVerify(c, header.Token); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	if err := p.planAccess.UpdateVerifyByPlanID(c, header.Token, uri.PlanID); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	data, err := p.planService.DeletePlan(c, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, data, "delete success!")
}

// CreateWorkout 創建訓練
// @Summary 創建訓練
// @Description 創建訓練
// @Tags Plan
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body validator.CreateWorkoutBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Workout} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /plan/{plan_id}/workout [POST]
func (p *Plan) CreateWorkout(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.PlanIDUri
	var body validator.CreateWorkoutBody
	if err := c.ShouldBindHeader(&header); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.trainerAccess.StatusVerify(c, header.Token); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	if err := p.workoutAccess.CreateVerifyByPlanID(c, header.Token, uri.PlanID); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	data, err := p.workoutService.CreateWorkout(c, uri.PlanID, body.Name)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, data, "create success!")
}

// GetWorkouts 取得計畫內的訓練列表
// @Summary  取得計畫內的訓練列表
// @Description  取得計畫內的訓練列表
// @Tags Plan
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=[]dto.Workout} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /plan/{plan_id}/workouts [GET]
func (p *Plan) GetWorkouts(c *gin.Context) {
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := p.workoutService.GetWorkoutsByPlanID(c, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, workouts, "get success!")
}
