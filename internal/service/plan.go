package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/plandto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type plan struct {
	planRepo repository.Plan
	courseRepo repository.Course
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewPlan(planRepo repository.Plan, courseRepo repository.Course, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &plan{planRepo: planRepo, courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (p *plan) CreatePlanByToken(c *gin.Context, token string, courseID int64, name string) (*plandto.Plan, errcode.Error) {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, p.errHandler.InvalidToken()
	}
	isExist, err := p.courseRepo.CheckCourseExistByIDAndUID(courseID, uid)
	if err != nil {
		return nil, p.errHandler.SystemError()
	}
	if !isExist {
		return nil, p.errHandler.PermissionDenied()
	}
	return p.CreatePlan(c, courseID, name)
}

func (p *plan) CreatePlan(c *gin.Context, courseID int64, name string) (*plandto.Plan, errcode.Error) {
	planID, err := p.planRepo.CreatePlan(courseID, name)
	if err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	var plan plandto.Plan
	if err := p.planRepo.FindPlanByID(planID, &plan); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plan, nil
}

func (p *plan) UpdatePlanByToken(c *gin.Context, token string, planID int64, name string) (*plandto.Plan, errcode.Error) {
	if err := p.checkPlanOwnerByPlanID(c, token, planID); err != nil {
		return nil, err
	}
	return p.UpdatePlan(c, planID, name)
}

func (p *plan) UpdatePlan(c *gin.Context, planID int64, name string) (*plandto.Plan, errcode.Error) {
	if err := p.checkPlanEditableByPlanID(c, planID); err != nil {
		return nil, err
	}
	if err := p.planRepo.UpdatePlanByID(planID, name); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	var plan plandto.Plan
	if err := p.planRepo.FindPlanByID(planID, &plan); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plan, nil
}

func (p *plan) GetPlansByCourseID(c *gin.Context, courseID int64) ([]*plandto.Plan, errcode.Error) {
	datas, err := p.planRepo.FindPlansByCourseID(courseID)
	if err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	plans := make([]*plandto.Plan, 0)
	for _, data := range datas {
		plan := plandto.Plan{
			ID: data.ID,
			Name: data.Name,
			WorkoutCount: data.WorkoutCount,
		}
		plans = append(plans, &plan)
	}
	return plans, nil
}

func (p *plan) DeletePlanByToken(c *gin.Context, token string, planID int64) (*plandto.PlanID, errcode.Error) {
	if err := p.checkPlanOwnerByPlanID(c, token, planID); err != nil {
		return nil, err
	}
	return p.DeletePlan(c, planID)
}

func (p *plan) DeletePlan(c *gin.Context, planID int64) (*plandto.PlanID, errcode.Error) {
	if err := p.checkPlanEditableByPlanID(c, planID); err != nil {
		return nil, err
	}
	if err := p.planRepo.DeletePlanByID(planID); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plandto.PlanID{ID: planID}, nil
}

func (p *plan) checkPlanOwnerByPlanID(c *gin.Context, token string, planID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	ownerID, err := p.planRepo.FindPlanOwnerByID(planID)
	if err != nil {
		p.logger.Set(c, handler.Error, "Plan", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if ownerID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *plan) checkPlanEditableByPlanID(c *gin.Context, planID int64) errcode.Error {
	status, err := p.courseRepo.FindCourseStatusByPlanID(planID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if !(status == 1 || status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}