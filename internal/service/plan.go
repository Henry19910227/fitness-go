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
	planRepo          repository.Plan
	courseRepo        repository.Course
	planStatisticRepo repository.UserPlanStatistic
	logger            handler.Logger
	jwtTool           tool.JWT
	errHandler        errcode.Handler
}

func NewPlan(planRepo repository.Plan, courseRepo repository.Course, planStatisticRepo repository.UserPlanStatistic, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Plan {
	return &plan{planRepo: planRepo, courseRepo: courseRepo, planStatisticRepo: planStatisticRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
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
		return nil, p.errHandler.Set(c, "plan repo", err)
	}
	plans := make([]*dto.Plan, 0)
	for _, data := range datas {
		plan := dto.Plan{
			ID:           data.ID,
			Name:         data.Name,
			WorkoutCount: data.WorkoutCount,
		}
		plans = append(plans, &plan)
	}
	return plans, nil
}

func (p *plan) GetPlanProductsByCourseID(c *gin.Context, userID int64, courseID int64) ([]*dto.PlanProduct, errcode.Error) {
	planDatas, err := p.planRepo.FindPlansByCourseID(courseID)
	if err != nil {
		return nil, p.errHandler.Set(c, "plan repo", err)
	}
	plans := make([]*dto.PlanProduct, 0)
	for _, planData := range planDatas {
		plan := dto.PlanProduct{
			ID:           planData.ID,
			Name:         planData.Name,
			WorkoutCount: planData.WorkoutCount,
		}
		plans = append(plans, &plan)
	}
	return plans, nil
}

func (p *plan) GetPlanAssets(c *gin.Context, userID int64, courseID int64) ([]*dto.PlanAsset, errcode.Error) {
	planDatas, err := p.planRepo.FindPlanAssets(userID, courseID)
	if err != nil {
		return nil, p.errHandler.Set(c, "plan repo", err)
	}
	plans := make([]*dto.PlanAsset, 0)
	for _, planData := range planDatas {
		plan := dto.PlanAsset{
			ID:                 planData.ID,
			Name:               planData.Name,
			WorkoutCount:       planData.WorkoutCount,
			FinishWorkoutCount: planData.FinishWorkoutCount,
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
