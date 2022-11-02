package card

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/card"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/card"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository card.Repository
}

func New(repository card.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *model.Table) (err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	err = s.repository.Create(item)
	return err
}
