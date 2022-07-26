package order_subscribe_plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/order_subscribe_plan"
	"gorm.io/gorm"
)

type service struct {
	repository order_subscribe_plan.Repository
}

func New(repository order_subscribe_plan.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *model.Table) (err error) {
	err = s.repository.Create(item)
	return err
}
