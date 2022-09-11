package min_duration_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/min_duration_record"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
