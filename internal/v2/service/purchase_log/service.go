package purchase_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/purchase_log"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/purchase_log"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository purchase_log.Repository
}

func New(repository purchase_log.Repository) Service {
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
