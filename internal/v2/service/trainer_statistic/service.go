package trainer_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
	"gorm.io/gorm"

	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer_statistic"
)

type service struct {
	repository trainer_statistic.Repository
}

func New(repository trainer_statistic.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) StatisticReviewScore(input *model.StatisticReviewScoreInput) (err error) {
	err = s.repository.StatisticReviewScore(input)
	return err
}

func (s *service) StatisticStudentCount() (err error) {
	err = s.repository.StatisticStudentCount()
	return err
}
