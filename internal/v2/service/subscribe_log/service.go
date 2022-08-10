package subscribe_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/subscribe_log"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository subscribe_log.Repository
}

func New(repository subscribe_log.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.CreateOrUpdate(item)
	return id, err
}
