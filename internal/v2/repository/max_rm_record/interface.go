package max_rm_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_rm_record"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}

