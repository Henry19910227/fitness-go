package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"strings"
)

type course struct {
	Base
	courseRepo repository.Course
	trainerRepo repository.Trainer
	uploader  handler.Uploader
	resHandler handler.Resource
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course,
	trainerRepo repository.Trainer,
	uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, trainerRepo: trainerRepo, uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (cs *course) CreateCourseByToken(c *gin.Context, token string, param *coursedto.CreateCourseParam) (*coursedto.Course, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	return cs.CreateCourse(c, uid, param)
}

func (cs *course) CreateCourse(c *gin.Context, uid int64, param *coursedto.CreateCourseParam) (*coursedto.Course, errcode.Error) {
	var courseID int64
	var err error
	if param.ScheduleType == 1 {
		courseID, err = cs.courseRepo.CreateSingleWorkoutCourse(uid, &model.CreateCourseParam{
			Name: param.Name,
			Level: param.Level,
			Category: param.Category,
		})
	} else {
		courseID, err = cs.courseRepo.CreateCourse(uid, &model.CreateCourseParam{
			Name: param.Name,
			Level: param.Level,
			Category: param.Category,
		})
	}
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	var course coursedto.Course
	if err := cs.courseRepo.FindCourseByID(courseID, &course); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &course, nil
}

func (cs *course) UpdateCourse(c *gin.Context, courseID int64, param *coursedto.UpdateCourseParam) (*coursedto.CourseDetail, errcode.Error) {
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		Category: param.Category,
		SaleID: param.SaleID,
		Name: param.Name,
		Intro: param.Intro,
		Food: param.Food,
		Level: param.Level,
		Suit: param.Suit,
		Equipment: param.Equipment,
		Place: param.Place,
		TrainTarget: param.TrainTarget,
		BodyTarget: param.BodyTarget,
		Notice: param.Notice,
	}); err != nil {
		if strings.Contains(err.Error(), "1452") {
			return nil, cs.errHandler.DataNotFound()
		}
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) DeleteCourse(c *gin.Context, courseID int64) (*coursedto.CourseID, errcode.Error) {
	if err := cs.courseRepo.DeleteCourseByID(courseID); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &coursedto.CourseID{ID: courseID}, nil
}

func (cs *course) GetCourseSummariesByToken(c *gin.Context, token string, status *int) ([]*coursedto.CourseSummary, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	return cs.GetCourseSummariesByUID(c, uid, status)
}

func (cs *course) GetCourseSummariesByUID(c *gin.Context, uid int64, status *int) ([]*coursedto.CourseSummary, errcode.Error) {
	Entities, err := cs.courseRepo.FindCourseSummariesByUserID(uid, status)
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	courses := make([]*coursedto.CourseSummary, 0)
	for _, entity := range Entities {
		course := coursedto.CourseSummary{
			ID:           entity.ID,
			CourseStatus: entity.CourseStatus,
			Category:     entity.Category,
			ScheduleType: entity.ScheduleType,
			SaleType:     entity.SaleType,
			Name:         entity.Name,
			Cover:        entity.Cover,
			Level:        entity.Level,
			PlanCount:    entity.PlanCount,
			WorkoutCount: entity.WorkoutCount,
		}
		course.Trainer.UserID = entity.Trainer.UserID
		course.Trainer.Nickname = entity.Trainer.Nickname
		course.Trainer.Avatar = entity.Trainer.Avatar
		courses = append(courses, &course)
	}
	return courses, nil
}

func (cs *course) GetCourseDetailByCourseID(c *gin.Context, courseID int64) (*coursedto.CourseDetail, errcode.Error) {
	entity, err := cs.courseRepo.FindCourseDetailByCourseID(courseID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, cs.errHandler.DataNotFound()
		}
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	course := coursedto.CourseDetail{
		ID:           entity.ID,
		CourseStatus: entity.CourseStatus,
		Category:     entity.Category,
		ScheduleType: entity.ScheduleType,
		SaleType:     entity.SaleType,
		Name:         entity.Name,
		Cover:        entity.Cover,
		Intro:        entity.Intro,
		Food:         entity.Food,
		Level:        entity.Level,
		Suit:         entity.Suit,
		Equipment:    entity.Equipment,
		Place:        entity.Place,
		TrainTarget:  entity.TrainTarget,
		BodyTarget:   entity.BodyTarget,
		Notice:       entity.Notice,
		PlanCount:    entity.PlanCount,
		WorkoutCount: entity.WorkoutCount,
		CreateAt:     entity.CreateAt,
		UpdateAt:     entity.UpdateAt,
	}
	course.Restricted = 0
	course.Trainer.UserID = entity.Trainer.UserID
	course.Trainer.Nickname = entity.Trainer.Nickname
	course.Trainer.Avatar = entity.Trainer.Avatar
	return &course, nil
}

func (cs *course) UploadCourseCoverByID(c *gin.Context, courseID int64, param *coursedto.UploadCourseCoverParam) (*coursedto.CourseCover, errcode.Error) {
	//上傳照片
	newImageNamed, err := cs.uploader.UploadCourseCover(param.File, param.CoverNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, cs.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, cs.errHandler.FileSizeError()
		}
		cs.logger.Set(c, handler.Error, "Resource Handler", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//查詢課表封面
	var course struct{ Cover string `gorm:"column:cover"`}
	if err := cs.courseRepo.FindCourseByID(courseID, &course); err != nil {
		cs.logger.Set(c, handler.Error, "Course Repo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//修改課表資訊
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		Cover: &newImageNamed,
	}); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//刪除舊照片
	if len(course.Cover) > 0 {
		if err := cs.resHandler.DeleteCourseCover(course.Cover); err != nil {
			cs.logger.Set(c, handler.Error, "ResHandler", cs.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &coursedto.CourseCover{Cover: newImageNamed}, nil
}

