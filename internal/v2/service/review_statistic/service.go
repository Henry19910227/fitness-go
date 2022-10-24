package review_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review_statistic"
	"gorm.io/gorm"
)

type service struct {
	repository review_statistic.Repository
}

func New(repository review_statistic.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Statistic(input *model.StatisticInput) (err error) {
	err = s.repository.Statistic(input)
	return err
}
