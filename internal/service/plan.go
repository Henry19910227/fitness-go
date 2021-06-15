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

func (p *plan) CreatePlanByToken(c *gin.Context, token string, courseID int64, name string) (*plandto.PlanID, errcode.Error) {
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

func (p *plan) CreatePlan(c *gin.Context, courseID int64, name string) (*plandto.PlanID, errcode.Error) {
	planID, err := p.planRepo.CreatePlan(courseID, name)
	if err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plandto.PlanID{ID: planID}, nil
}
