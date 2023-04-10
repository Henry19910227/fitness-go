package course_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course_training_monthly_statistic"
	"gorm.io/gorm"
)

type service struct {
	repository course_training_monthly_statistic.Repository
}

func New(repository course_training_monthly_statistic.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	if err != nil {
		return output, err
	}
	return output, err
}

func (s *service) List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.List(input)
	if err != nil {
		return output, page, err
	}
	page = &paging.Output{}
	page.TotalCount = int(amount)
	page.Page = input.Page
	page.Size = input.Size
	if input.Size != nil {
		page.TotalPage = util.PointerInt(util.Pagination(int(amount), *input.Size))
	}
	return output, page, err
}

func (s *service) Statistic(input *model.StatisticInput) (err error) {
	err = s.repository.Statistic(input)
	return err
}
