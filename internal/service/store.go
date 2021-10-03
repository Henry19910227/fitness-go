package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/dto/saledto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type store struct {
	Base
	courseRepo repository.Course
	trainerRepo repository.Trainer
	errHandler errcode.Handler
}

func NewStore(courseRepo repository.Course, trainerRepo repository.Trainer, errHandler errcode.Handler) Store {
	return &store{courseRepo: courseRepo, trainerRepo: trainerRepo, errHandler: errHandler}
}

func (s *store) GetHomePage(c *gin.Context) (*dto.StoreHomePage, errcode.Error) {
	latestCourses, err := s.getLatestCourseSummaries()
	if err != nil {
		return nil, s.errHandler.Set(c, "store", err)
	}
	latestTrainers, err := s.getLatestTrainerSummaries()
	if err != nil {
		return nil, s.errHandler.Set(c, "store", err)
	}
	return &dto.StoreHomePage{LatestTrainers: latestTrainers,
		PopularTrainers: latestTrainers,
		LatestCourses: latestCourses,
		PopularCourses: latestCourses}, nil
}

func (s *store) GetCourseProduct(c *gin.Context, page, size int) ([]*dto.CourseSummary, errcode.Error) {
	offset, limit := s.GetPagingIndex(page, size)
	var status = global.Sale
	entities, err := s.courseRepo.FindCourseSummaries(&model.FindCourseSummariesParam{
		Status: &status,
	}, &model.OrderBy{
		Field:     "courses.update_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, s.errHandler.Set(c, "store", err)
	}
	return parserCourses(entities), nil
}


func (s *store) getLatestTrainerSummaries() ([]*dto.TrainerSummary, error) {
	trainers := make([]*dto.TrainerSummary, 0)
	var trainerStatus = global.TrainerActivity
	if err := s.trainerRepo.FindTrainers(&trainers, &trainerStatus, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  5,
	}); err != nil {
		return nil, err
	}
	return trainers, nil
}

func (s *store) getLatestCourseSummaries() ([]*dto.CourseSummary, error) {
	var status = global.Sale
	entities, err := s.courseRepo.FindCourseSummaries(&model.FindCourseSummariesParam{
		Status: &status,
	}, &model.OrderBy{
		Field:     "courses.update_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  5,
	})
	if err != nil {
		return nil, err
	}
	return parserCourses(entities), nil
}

func parserCourses(entities []*model.CourseSummaryEntity) []*dto.CourseSummary {
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
		}
		course.Trainer = trainer
		if entity.Sale.ID != 0 {
			sale := &saledto.SaleItem{
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
	return courses
}