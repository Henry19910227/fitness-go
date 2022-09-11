package min_duration_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/min_duration_record"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/min_duration_record"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository min_duration_record.Repository
}

func New(repository min_duration_record.Repository) Service {
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