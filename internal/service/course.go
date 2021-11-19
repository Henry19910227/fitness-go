package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type course struct {
	Base
	courseRepo repository.Course
	trainerRepo repository.Trainer
	planRepo  repository.Plan
	uploader  handler.Uploader
	resHandler handler.Resource
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course,
	trainerRepo repository.Trainer,
	planRepo  repository.Plan,
	uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, trainerRepo: trainerRepo, planRepo: planRepo, uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (cs *course) CreateCourseByToken(c *gin.Context, token string, param *dto.CreateCourseParam) (*dto.Course, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	return cs.CreateCourse(c, uid, param)
}

func (cs *course) CreateCourse(c *gin.Context, uid int64, param *dto.CreateCourseParam) (*dto.Course, errcode.Error) {
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
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) UpdateCourse(c *gin.Context, courseID int64, param *dto.UpdateCourseParam) (*dto.Course, errcode.Error) {
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
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) DeleteCourse(c *gin.Context, courseID int64) (*dto.CourseID, errcode.Error) {
	if err := cs.courseRepo.DeleteCourseByID(courseID); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &dto.CourseID{ID: courseID}, nil
}

func (cs *course) GetCourseSummariesByUID(c *gin.Context, uid int64, status *int) ([]*dto.CourseSummary, errcode.Error) {
	entities, err := cs.courseRepo.FindCourseSummaries(&model.FindCourseSummariesParam{
		UID: &uid,
		Status: status,
	}, nil, nil)
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	courses := make([]*dto.CourseSummary, 0)
	for _, entity := range entities {
		course := dto.CourseSummary{
			ID:           entity.ID,
			CourseStatus: entity.CourseStatus,
			Category:     entity.Category,
			ScheduleType: entity.ScheduleType,
			Name:         entity.Name,
			Cover:        entity.Cover,
			Level:        entity.Level,
			PlanCount:    entity.PlanCount,
			WorkoutCount: entity.WorkoutCount,
		}
		trainer := &dto.TrainerSummary{
			UserID: entity.Trainer.UserID,
			Nickname: entity.Trainer.Nickname,
			Avatar: entity.Trainer.Avatar,
			Skill: entity.Trainer.Skill,
		}
		course.Trainer = trainer
		if entity.Sale.ID != 0 {
			sale := &dto.SaleItem{
				ID: entity.Sale.ID,
				Type: entity.Sale.Type,
				Name: entity.Sale.Name,
				Twd: entity.Sale.Twd,
				Identifier: entity.Sale.Identifier,
			}
			course.Sale = sale
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (cs *course) GetCourseDetailByCourseID(c *gin.Context, courseID int64) (*dto.Course, errcode.Error) {
	entity, err := cs.courseRepo.FindCourseDetailByCourseID(courseID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, cs.errHandler.DataNotFound()
		}
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	course := dto.Course{
		ID:           entity.ID,
		CourseStatus: entity.CourseStatus,
		Category:     entity.Category,
		ScheduleType: entity.ScheduleType,
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
	trainer := &dto.TrainerSummary{
		UserID: entity.Trainer.UserID,
		Nickname: entity.Trainer.Nickname,
		Avatar: entity.Trainer.Avatar,
		Skill: entity.Trainer.Skill,
	}
	course.Trainer = trainer
	if entity.Sale.ID != 0 {
		sale := &dto.SaleItem{
			ID: entity.Sale.ID,
			Type: entity.Sale.Type,
			Name: entity.Sale.Name,
			Twd: entity.Sale.Twd,
			Identifier: entity.Sale.Identifier,
		}
		course.Sale = sale
	}
	course.Restricted = 0
	return &course, nil
}

func (cs *course) GetCourseProductByCourseID(c *gin.Context, courseID int64) (*dto.CourseProduct, errcode.Error) {
	course, err := cs.parserCourseProduct(courseID)
	if err != nil {
		return course, cs.errHandler.Set(c, "course repo", err)
	}
	course.AllowAccess = 0
	return course, nil
}

func (cs *course) GetCourseOverviewByCourseID(c *gin.Context, courseID int64) (*dto.CourseProduct, errcode.Error) {
	course, err := cs.parserCourseProduct(courseID)
	if err != nil {
		return course, cs.errHandler.Set(c, "course repo", err)
	}
	course.AllowAccess = 1
	return course, nil
}

func (cs *course) GetCourseProductSummaries(c *gin.Context, param *dto.GetCourseProductSummariesParam, page, size int) ([]*dto.CourseProductSummary, errcode.Error) {
	var field = string(global.UpdateAt)
	offset, limit := cs.GetPagingIndex(page, size)
	datas, err := cs.courseRepo.FindCourseProductSummaries(model.FindCourseProductSummariesParam{
		Name: param.Name,
		Score: param.Score,
		Level: param.Level,
		Category: param.Category,
		Suit: param.Suit,
		Equipment: param.Equipment,
		Place: param.Place,
		TrainTarget: param.TrainTarget,
		BodyTarget: param.BodyTarget,
		SaleType: param.SaleType,
		TrainerSex: param.TrainerSex,
		TrainerSkill: param.TrainerSkill,
	}, &model.OrderBy{
		Field:     field,
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, cs.errHandler.Set(c, "store", err)
	}
	return parserCourses(datas), nil
}

func (cs *course) UploadCourseCoverByID(c *gin.Context, courseID int64, param *dto.UploadCourseCoverParam) (*dto.CourseCover, errcode.Error) {
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
	return &dto.CourseCover{Cover: newImageNamed}, nil
}

func (cs *course) CourseSubmit(c *gin.Context, courseID int64) errcode.Error {
	//驗證課表填寫完整性
	entity, err := cs.courseRepo.FindCourseDetailByCourseID(courseID)
	if err != nil {
		return cs.errHandler.Set(c, "course repo", err)
	}
	if err := cs.VerifyCourse(entity); err != nil {
		return cs.errHandler.Set(c, "verify course", err)
	}
	//送審課表(測試暫時將課表狀態改為"銷售中")
	var courseStatus = global.Sale
	if err := cs.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		CourseStatus: &courseStatus,
	}); err != nil {
		return cs.errHandler.Set(c, "course repo", err)
	}
	return nil
}

func (cs *course) GetCourseStatus(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error) {
	course := struct {
		CourseStatus int `json:"course_status"`
	}{}
	if err := cs.courseRepo.FindCourseByID(courseID, &course); err != nil {
		return 0, cs.errHandler.Set(c, "course repo", err)
	}
	return global.CourseStatus(course.CourseStatus), nil
}

func (cs *course) VerifyCourse(course *model.CourseDetailEntity) error {
	if course.Sale.ID == 0 {
		return errors.New(strconv.Itoa(errcode.UpdateError))
	}
	if course.Trainer.UserID == 0 {
		return errors.New(strconv.Itoa(errcode.UpdateError))
	}
	return nil
}

func (cs *course) parserCourseProduct(courseID int64) (*dto.CourseProduct, error) {
	//查詢課表詳情
	courseItem, err := cs.courseRepo.FindCourseProduct(courseID)
	if err != nil {
		return nil, err
	}
	course := dto.CourseProduct{
		ID:           courseItem.ID,
		CourseStatus: courseItem.CourseStatus,
		Category:     courseItem.Category,
		ScheduleType: courseItem.ScheduleType,
		Name:         courseItem.Name,
		Cover:        courseItem.Cover,
		Intro:        courseItem.Intro,
		Food:         courseItem.Food,
		Level:        courseItem.Level,
		Suit:         courseItem.Suit,
		Equipment:    courseItem.Equipment,
		Place:        courseItem.Place,
		TrainTarget:  courseItem.TrainTarget,
		BodyTarget:   courseItem.BodyTarget,
		Notice:       courseItem.Notice,
		PlanCount:    courseItem.PlanCount,
		WorkoutCount: courseItem.WorkoutCount,
	}
	//配置教練資訊
	trainer := &dto.TrainerSummary{
		UserID:   courseItem.Trainer.UserID,
		Nickname: courseItem.Trainer.Nickname,
		Avatar:   courseItem.Trainer.Avatar,
		Skill:    courseItem.Trainer.Skill,
	}
	course.Trainer = trainer
	//配置銷售資訊
	if courseItem.Sale.ID != 0 {
		sale := &dto.SaleItem{
			ID:         courseItem.Sale.ID,
			Type:       courseItem.Sale.Type,
			Name:       courseItem.Sale.Name,
			Twd:        courseItem.Sale.Twd,
			Identifier: courseItem.Sale.Identifier,
		}
		course.Sale = sale
	}
	review := dto.ReviewStatistic{
		ScoreTotal: courseItem.Review.ScoreTotal,
		Amount: courseItem.Review.Amount,
		FiveTotal: courseItem.Review.FiveTotal,
		FourTotal: courseItem.Review.FourTotal,
		ThreeTotal: courseItem.Review.ThreeTotal,
		TwoTotal: courseItem.Review.TwoTotal,
		OneTotal: courseItem.Review.OneTotal,
		UpdateAt: courseItem.Review.UpdateAt,
	}
	course.Review = review
	return &course, nil
}


