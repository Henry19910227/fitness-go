package access

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"strings"
)

type course struct {
	courseRepo  repository.Course
	trainerRepo repository.Trainer
	logger      handler.Logger
	jwtTool     tool.JWT
	errHandler  errcode.Handler
}

func NewCourse(courseRepo repository.Course,
	trainerRepo repository.Trainer,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo,
		trainerRepo: trainerRepo,
		logger:      logger,
		jwtTool:     jwtTool,
		errHandler:  errHandler}
}

func (p *course) CreateVerify(c *gin.Context, token string) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	var trainer struct {
		UserID        int64 `gorm:"column:user_id"`
		TrainerStatus int   `gorm:"column:trainer_status"`
	}
	if err := p.trainerRepo.FindTrainerEntity(uid, &trainer); err != nil {
		p.logger.Set(c, handler.Error, "TrainerRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if trainer.UserID == 0 {
		return p.errHandler.PermissionDenied()
	}
	//教練審核中狀態
	if trainer.TrainerStatus == 2 {
		amount, err := p.courseRepo.FindCourseAmountByUserID(trainer.UserID)
		if err != nil {
			p.logger.Set(c, handler.Error, "TrainerRepo", p.errHandler.SystemError().Code(), err.Error())
			return p.errHandler.SystemError()
		}
		if amount > 1 {
			return p.errHandler.PermissionDenied()
		}
	}
	//教練停權狀態
	if trainer.TrainerStatus == 3 {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *course) UpdateVerifyByCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int   `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByID(nil, courseID, &course); err != nil {
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

func (p *course) OwnVerifyByTokenAndCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int   `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByID(nil, courseID, &course); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return p.errHandler.DataNotFound()
		}
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if course.UserID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}
