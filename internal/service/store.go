package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type store struct {
	Base
	courseRepo  repository.Course
	trainerRepo repository.Trainer
	reviewRepo  repository.Review
	errHandler  errcode.Handler
}

func NewStore(courseRepo repository.Course, trainerRepo repository.Trainer, reviewRepo repository.Review, errHandler errcode.Handler) Store {
	return &store{courseRepo: courseRepo, trainerRepo: trainerRepo, reviewRepo: reviewRepo, errHandler: errHandler}
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
		LatestCourses:   latestCourses,
		PopularCourses:  latestCourses}, nil
}

func (s *store) getLatestTrainerSummaries() ([]*dto.TrainerSummary, error) {
	trainers := make([]*dto.TrainerSummary, 0)
	var trainerStatus = global.TrainerActivity
	if err := s.trainerRepo.FindTrainerEntities(&trainers, &trainerStatus, &model.OrderBy{
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

func (s *store) getLatestCourseSummaries() ([]*dto.CourseProductSummary, error) {
	entities, err := s.courseRepo.FindCourseProductSummaries(model.FindCourseProductSummariesParam{}, &model.OrderBy{
		Field:     "update_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  5,
	})
	if err != nil {
		return nil, err
	}
	return parserCourseProductSummaries(entities), nil
}

func parserCourseProductSummaries(datas []*model.CourseProductSummary) []*dto.CourseProductSummary {
	courses := make([]*dto.CourseProductSummary, 0)
	for _, data := range datas {
		course := dto.CourseProductSummary{
			ID:           data.ID,
			SaleType:     data.SaleType,
			CourseStatus: data.CourseStatus,
			Category:     data.Category,
			ScheduleType: data.ScheduleType,
			Name:         data.Name,
			Cover:        data.Cover,
			Level:        data.Level,
			PlanCount:    data.PlanCount,
			WorkoutCount: data.WorkoutCount,
		}
		if data.Trainer != nil {
			course.Trainer = &dto.TrainerSummary{
				UserID:   data.Trainer.UserID,
				Nickname: data.Trainer.Nickname,
				Avatar:   data.Trainer.Avatar,
				Skill:    data.Trainer.Skill,
			}
		}
		course.Review = dto.ReviewStatisticSummary{
			ScoreTotal: data.Review.ScoreTotal,
			Amount:     data.Review.Amount,
		}
		if data.Sale != nil {
			sale := &dto.SaleItem{
				ID:   data.Sale.ID,
				Type: data.Sale.Type,
				Name: data.Sale.Name,
			}
			course.Sale = sale
			if data.Sale.ProductLabel != nil {
				course.Sale.Twd = data.Sale.ProductLabel.Twd
				course.Sale.ProductID = data.Sale.ProductLabel.ProductID
			}
		}
		courses = append(courses, &course)
	}
	return courses
}
