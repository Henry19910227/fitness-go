package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type permissions struct {
	Base
	courseRepo repository.Course
	trainerRepo repository.Trainer
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewPermissions(courseRepo repository.Course,
	trainerRepo repository.Trainer,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Permissions {
	return &permissions{courseRepo: courseRepo,
		trainerRepo: trainerRepo,
		logger: logger,
		jwtTool: jwtTool,
		errHandler: errHandler}
}

func (p *permissions) CheckTrainerValidByUID(c *gin.Context, token string) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	var trainer struct{TrainerStatus int `gorm:"column:trainer_status"`}
	if err := p.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil{
		return p.errHandler.PermissionDenied()
	}
	if trainer.TrainerStatus == 3 {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckCourseOwnerByCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	ownerID, err := p.courseRepo.FindCourseOwnerByID(courseID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if ownerID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckPlanOwnerByPlanID(c *gin.Context, token string, planID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	ownerID, err := p.courseRepo.FindCourseOwnerByPlanID(planID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if ownerID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckWorkoutOwnerByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	ownerID, err := p.courseRepo.FindCourseOwnerByWorkoutID(workoutID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if ownerID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckActionOwnerByActionID(c *gin.Context, token string, actionID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	ownerID, err := p.courseRepo.FindCourseOwnerByActionID(actionID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if ownerID != uid {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckCourseEditableByCourseID(c *gin.Context, courseID int64) errcode.Error {
	status, err := p.courseRepo.FindCourseStatusByID(courseID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if !(status == 1 || status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckPlanEditableByPlanID(c *gin.Context, planID int64) errcode.Error {
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

func (p *permissions) CheckWorkoutEditableByWorkoutID(c *gin.Context, workoutID int64) errcode.Error {
	status, err := p.courseRepo.FindCourseStatusByWorkoutID(workoutID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if !(status == 1 || status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) CheckActionEditableByActionID(c *gin.Context, actionID int64) errcode.Error {
	status, err := p.courseRepo.FindCourseStatusByActionID(actionID)
	if err != nil {
		p.logger.Set(c, handler.Error, "CourseRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if !(status == 1 || status == 4) {
		return p.errHandler.PermissionDenied()
	}
	return nil
}
