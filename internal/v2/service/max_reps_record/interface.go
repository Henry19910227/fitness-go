package max_reps_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_reps_record"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
