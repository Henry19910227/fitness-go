package workout_set_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set_order"
	"gorm.io/gorm"
)

type service struct {
	repository workout_set_order.Repository
}

func New(repository workout_set_order.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	err = s.repository.Create(items)
	return err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}