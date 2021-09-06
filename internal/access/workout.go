package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type workout struct {
	courseRepo repository.Course
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewWorkout(courseRepo repository.Course,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Workout {
	return &workout{courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (p *workout) CreateVerifyByPlanID(c *gin.Context, uid int64, planID int64) errcode.Error {
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
		ScheduleType int `gorm:"column:schedule_type"`
	}{}
	if err := p.courseRepo.FindCourseByPlanID(planID, &course); err != nil {
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

func (p *workout) UpdateVerifyByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByWorkoutID(workoutID, &course); err != nil {
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
