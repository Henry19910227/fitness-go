package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
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

func (p *plan) CreatePlan(c *gin.Context, courseID int64, name string) (*dto.Plan, errcode.Error) {
	planID, err := p.planRepo.CreatePlan(courseID, name)
	if err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	var plan dto.Plan
	if err := p.planRepo.FindPlanByID(planID, &plan); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plan, nil
}

func (p *plan) UpdatePlan(c *gin.Context, planID int64, name string) (*dto.Plan, errcode.Error) {
	if err := p.planRepo.UpdatePlanByID(planID, name); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	var plan dto.Plan
	if err := p.planRepo.FindPlanByID(planID, &plan); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &plan, nil
}

func (p *plan) GetPlansByCourseID(c *gin.Context, courseID int64) ([]*dto.Plan, errcode.Error) {
	datas, err := p.planRepo.FindPlansByCourseID(courseID)
	if err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	plans := make([]*dto.Plan, 0)
	for _, data := range datas {
		plan := dto.Plan{
			ID: data.ID,
			Name: data.Name,
			WorkoutCount: data.WorkoutCount,
		}
		plans = append(plans, &plan)
	}
	return plans, nil
}

func (p *plan) DeletePlan(c *gin.Context, planID int64) (*dto.PlanID, errcode.Error) {
	if err := p.planRepo.DeletePlanByID(planID); err != nil {
		p.logger.Set(c, handler.Error, "PlanRepo", p.errHandler.SystemError().Code(), err.Error())
		return nil, p.errHandler.SystemError()
	}
	return &dto.PlanID{ID: planID}, nil
}

func (p *plan) GetPlanStatus(c *gin.Context, planID int64) (global.CourseStatus, errcode.Error) {
	course := struct {
		CourseStatus int `json:"course_status"`
	}{}
	if err := p.courseRepo.FindCourseByPlanID(planID, &course); err != nil {
		return 0, p.errHandler.Set(c, "course repo", err)
	}
	return global.CourseStatus(course.CourseStatus), nil
}