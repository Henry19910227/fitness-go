package max_weight_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_weight_record"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/max_weight_record"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository max_weight_record.Repository
}

func New(repository max_weight_record.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.CreateOrUpdate(item)
	return id, err
}
