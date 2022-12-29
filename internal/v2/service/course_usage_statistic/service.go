package course_usage_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course_usage_statistic"
	"gorm.io/gorm"
)

type service struct {
	repository course_usage_statistic.Repository
}

func New(repository course_usage_statistic.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Statistic() (err error) {
	err = s.repository.Statistic()
	return err
}
