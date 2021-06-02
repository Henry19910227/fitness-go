package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type course struct {
	Base
	courseRepo repository.Course
	uploader  handler.Uploader
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course, uploader handler.Uploader, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, uploader: uploader, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
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
		ScheduleType: param.ScheduleType,
	})
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &coursedto.CreateResult{ID: courseID}, nil
}

func (cs *course) UpdateCourse(c *gin.Context, courseID int64, param *coursedto.UpdateCourseParam) (*coursedto.Course, errcode.Error) {
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		Category: param.Category,
		ScheduleType: param.ScheduleType,
		SaleType: param.SaleType,
		Price: param.Price,
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

func (cs *course) GetCourseByTokenAndCourseID(c *gin.Context, token string, courseID int64) (*coursedto.Course, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	course, e := cs.GetCourseByID(c, courseID)
	if e != nil {
		return nil, e
	}
	//驗證權限
	if course.UserID != uid {
		return nil, cs.errHandler.PermissionDenied()
	}
	return course, nil
}

func (cs *course) GetCourseByID(c *gin.Context, courseID int64) (*coursedto.Course, errcode.Error) {
	//取得課表資料
	var course coursedto.Course
	if err := cs.courseRepo.FindCourseByID(courseID, &course); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.DataNotFound()
		}
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
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
	//修改課表資訊
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		Cover: &newImageNamed,
	}); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &coursedto.CourseCover{Cover: newImageNamed}, nil
}

