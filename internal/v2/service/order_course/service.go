package order_course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/order_course"
	"gorm.io/gorm"
)

type service struct {
	repository order_course.Repository
}

func New(repository order_course.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *model.Table) (err error) {
	err = s.repository.Create(item)
	return err
}
