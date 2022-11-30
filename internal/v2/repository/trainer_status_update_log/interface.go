package trainer_status_update_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_status_update_log"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (id int64, err error)
	Creates(items []*model.Table) (err error)
}
