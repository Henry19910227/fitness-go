package trainer_album

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_album"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer_album"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository trainer_album.Repository
}

func New(repository trainer_album.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *model.Table) (id int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.Create(item)
	return id, err
}

func (s *service) Creates(items []*model.Table) (err error) {
	for _, item := range items {
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	err = s.repository.Creates(items)
	return err
}