package subscribe_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
