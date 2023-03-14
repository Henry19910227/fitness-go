package user_plan_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_plan_statistic"
	"gorm.io/gorm"
)

type service struct {
	repository user_plan_statistic.Repository
}

func New(repository user_plan_statistic.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Statistic(input *model.Statistic) (err error) {
	err = s.repository.Statistic(input)
	return err
}
