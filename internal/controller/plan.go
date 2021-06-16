package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Plan struct {
	Base
	planService service.Plan
}

func NewPlan(baseGroup *gin.RouterGroup, planService service.Plan, userMiddleware gin.HandlerFunc)  {
	plan := Plan{planService: planService}
	planGroup := baseGroup.Group("/plan")
	planGroup.Use(userMiddleware)
	planGroup.PATCH("/:plan_id", plan.UpdatePlan)
	planGroup.DELETE("/:plan_id", plan.DeletePlan)
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
	planData, err := p.planService.UpdatePlanByToken(c, header.Token, uri.PlanID, body.Name)
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
	data, err := p.planService.DeletePlanByToken(c, header.Token, uri.PlanID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, data, "delete success!")
}
