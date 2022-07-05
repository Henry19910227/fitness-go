package feedback_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/feedback_image"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository feedback_image.Repository
}

func New(repository feedback_image.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(items []*model.Table) (err error) {
	for _, item := range items{
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	err = s.repository.Create(items)
	return err
}
