package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type action struct {
	courseRepo repository.Course
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewAction(courseRepo repository.Course,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Action {
	return &action{courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (a *action) CreateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return a.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := a.courseRepo.FindCourseByID(courseID, &course); err != nil {
		a.logger.Set(c, handler.Error, "CourseRepo", a.errHandler.SystemError().Code(), err.Error())
		return a.errHandler.SystemError()
	}
	if course.UserID != uid {
		return a.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return a.errHandler.PermissionDenied()
	}
	return nil
}

func (a *action) UpdateVerifyByActionID(c *gin.Context, token string, actionID int64) errcode.Error {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return a.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := a.courseRepo.FindCourseByActionID(actionID, &course); err != nil {
		a.logger.Set(c, handler.Error, "CourseRepo", a.errHandler.SystemError().Code(), err.Error())
		return a.errHandler.SystemError()
	}
	if course.UserID != uid {
		return a.errHandler.PermissionDenied()
	}
	if !(course.Status == 1 || course.Status == 4) {
		return a.errHandler.PermissionDenied()
	}
	return nil
}
