package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review_image"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository review_image.Repository
}

func New(repository review_image.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
}

func (s *service) Create(items []*model.Table) (err error) {
	for _, item := range items {
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	err = s.repository.Create(items)
	return err
}
