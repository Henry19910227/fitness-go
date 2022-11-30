package course_status_update_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_status_update_log"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (id int64, err error)
	Creates(items []*model.Table) (err error)
}
