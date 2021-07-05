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

func (p *permissions) CourseValidationByCourseID(c *gin.Context, token string, courseID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	if err := p.checkTrainerStatusByUID(c, uid); err != nil {
		return err
	}
	if err := p.checkCourseEditableByCourseID(c, uid, courseID); err != nil {
		return err
	}
	return nil
}

func (p *permissions) CourseValidationByPlanID(c *gin.Context, token string, planID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	if err := p.checkTrainerStatusByUID(c, uid); err != nil {
		return err
	}
	if err := p.checkCourseEditableByPlanID(c, uid, planID); err != nil {
		return err
	}
	return nil
}

func (p *permissions) CourseValidationByWorkoutID(c *gin.Context, token string, workoutID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	if err := p.checkTrainerStatusByUID(c, uid); err != nil {
		return err
	}
	if err := p.checkCourseEditableByWorkoutID(c, uid, workoutID); err != nil {
		return err
	}
	return nil
}

func (p *permissions) CourseValidationByActionID(c *gin.Context, token string, actionID int64) errcode.Error {
	uid, err := p.jwtTool.GetIDByToken(token)
	if err != nil {
		return p.errHandler.InvalidToken()
	}
	if err := p.checkTrainerStatusByUID(c, uid); err != nil {
		return err
	}
	if err := p.checkCourseEditableByActionID(c, uid, actionID); err != nil {
		return err
	}
	return nil
}

func (p *permissions) checkTrainerStatusByUID(c *gin.Context, uid int64) errcode.Error {
	var trainer struct{TrainerStatus int `gorm:"column:trainer_status"`}
	if err := p.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil{
		p.logger.Set(c, handler.Error, "TrainerRepo", p.errHandler.SystemError().Code(), err.Error())
		return p.errHandler.SystemError()
	}
	if trainer.TrainerStatus == 3 {
		return p.errHandler.PermissionDenied()
	}
	return nil
}

func (p *permissions) checkCourseEditableByCourseID(c *gin.Context, uid int64, courseID int64) errcode.Error {
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByID(courseID, &course); err != nil {
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

func (p *permissions) checkCourseEditableByPlanID(c *gin.Context, uid int64, planID int64) errcode.Error {
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
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

func (p *permissions) checkCourseEditableByWorkoutID(c *gin.Context, uid int64, workoutID int64) errcode.Error {
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

func (p *permissions) checkCourseEditableByActionID(c *gin.Context, uid int64, actionID int64) errcode.Error {
	course := struct {
		UserID int64 `gorm:"column:user_id"`
		Status int `gorm:"column:course_status"`
	}{}
	if err := p.courseRepo.FindCourseByActionID(actionID, &course); err != nil {
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