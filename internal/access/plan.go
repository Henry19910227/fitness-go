package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type plan struct {
	courseRepo repository.Course
	logger     handler.Logger
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewPlan(courseRepo repository.Course,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Plan {
	return &plan{courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (p *plan) CreateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	course := struct {
		UserID       int64 `gorm:"column:user_id"`
		Status       int   `gorm:"column:course_status"`
		ScheduleType int   `gorm:"column:schedule_type"`
	}{}
	if err := p.courseRepo.FindCourseByID(nil, courseID, &course); err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if course.UserID != uid {
		return p.errHandler.PermissionDenied()
	}
	if course.ScheduleType == 1 {
		return p.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *plan) UpdateVerifyByPlanID(c *gin.Context, token string, planID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int   `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByPlanID(planID, &course); err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if course.UserID != uid {
		return p.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}
