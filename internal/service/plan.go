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
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, p.errHandler.InvalidToken()
	}
	isExist, err := p.planRepo.CheckPlanExistByUID(uid, planID)
	if err != nil {
		return nil, p.errHandler.SystemError()
	}
	if !isExist {
		return nil, p.errHandler.PermissionDenied()
	}
	return p.DeletePlan(c, planID)
}

func (p *plan) DeletePlan(c *gin.Context, planID int64) (*plandto.PlanID, errcode.Error) {
	if err := p.planRepo.DeletePlanByID(planID); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plandto.PlanID{ID: planID}, nil
}
