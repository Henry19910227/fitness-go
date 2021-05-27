package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type course struct {
	Base
	courseRepo repository.Course
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (cs *course) CreateCourseByToken(c *gin.Context, token string, param *coursedto.CreateCourseParam) (*coursedto.CreateResult, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	return cs.CreateCourse(c, uid, param)
}

func (cs *course) CreateCourse(c *gin.Context, uid int64, param *coursedto.CreateCourseParam) (*coursedto.CreateResult, errcode.Error) {
	courseID, err := cs.courseRepo.CreateCourse(uid, &model.CreateCourseParam{
		Name: param.Name,
		Level: param.Level,
		Category: param.Category,
		CategoryOther: param.CategoryOther,
		ScheduleType: param.ScheduleType,
	})
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &coursedto.CreateResult{ID: courseID}, nil
}

func (cs *course) UpdateCourse(c *gin.Context, courseID int64, param *coursedto.UpdateCourseParam) errcode.Error {
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		CourseStatus: param.CourseStatus,
		Category: param.Category,
		ScheduleType: param.ScheduleType,
		SaleType: param.SaleType,
		Price: param.Price,
		Name: param.Name,
		Image: param.Image,
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
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return cs.errHandler.SystemError()
	}
	return nil
}
