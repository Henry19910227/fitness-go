package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/access"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type Plan struct {
	Base
	planService      service.Plan
	workoutService   service.Workout
	workoutSetAccess access.WorkoutSet
}

func NewPlan(baseGroup *gin.RouterGroup,
	planService service.Plan,
	workoutService service.Workout,
	workoutSetAccess access.WorkoutSet,
	userMidd midd.User,
	courseMidd midd.Course) {
	plan := Plan{planService: planService,
		workoutService:   workoutService,
		workoutSetAccess: workoutSetAccess}

	baseGroup.PATCH("/plan/:plan_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		plan.UpdatePlan)

	baseGroup.DELETE("/plan/:plan_id",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		plan.DeletePlan)

	baseGroup.POST("/plan/:plan_id/workout",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		courseMidd.AdminAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		plan.CreateWorkout)

	baseGroup.GET("/plan/:plan_id/workouts",
		userMidd.TokenPermission([]global.Role{global.UserRole, global.AdminRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		plan.GetWorkouts)
}

// UpdatePlan 修改計畫
// @Summary 修改計畫 (API已經過時，更新為 /v2/trainer/plan/{plan_id}/workout [POST])
// @Description 修改計畫
// @Tags Plan_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body validator.UpdatePlanBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Plan} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /v1/plan/{plan_id} [PATCH]
func (p *Plan) UpdatePlan(c *gin.Context) {
	var uri validator.PlanIDUri
	var body validator.UpdatePlanBody
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
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
// @Summary 刪除計畫 (API已經過時，更新為 /v2/trainer/plan/{plan_id} [DELETE])
// @Description 刪除計畫
// @Tags Plan_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=dto.PlanID} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/plan/{plan_id} [DELETE]
func (p *Plan) DeletePlan(c *gin.Context) {
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
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
// @Summary 創建訓練 (API已經過時，更新為 /v2/trainer/plan/{plan_id}/workout [POST])
// @Description 創建訓練
// @Tags Plan_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Param json_body body validator.CreateWorkoutBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Workout} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /v1/plan/{plan_id}/workout [POST]
func (p *Plan) CreateWorkout(c *gin.Context) {
	var uri validator.PlanIDUri
	var body validator.CreateWorkoutBody
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//直接創建訓練
	if body.WorkoutTemplateID == nil {
		data, err := p.workoutService.CreateWorkout(c, uri.PlanID, body.Name)
		if err != nil {
			p.JSONErrorResponse(c, err)
			return
		}
		p.JSONSuccessResponse(c, data, "create success!")
		return
	}
	//使用訓練模板複製訓練
	uid, e := p.GetUID(c)
	if e != nil {
		p.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	if err := p.workoutSetAccess.CreateVerifyByWorkoutID(c, uid, *body.WorkoutTemplateID); err != nil {
		p.JSONValidatorErrorResponse(c, err.Msg())
		return
	}
	data, err := p.workoutService.CreateWorkoutByTemplate(c, uri.PlanID, body.Name, *body.WorkoutTemplateID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, data, "create success!")
}

// GetWorkouts 取得計畫內的訓練列表
// @Summary 取得計畫內的訓練列表 (API已經過時，更新為 /v2/trainer/plan/{plan_id}/workouts [GET])
// @Description  取得計畫內的訓練列表
// @Tags Plan_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param plan_id path int64 true "計畫id"
// @Success 200 {object} model.SuccessResult{data=[]dto.Workout} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/plan/{plan_id}/workouts [GET]
func (p *Plan) GetWorkouts(c *gin.Context) {
	var uri validator.PlanIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	workouts, err := p.workoutService.GetWorkouts(c, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, workouts, "get success!")
}
