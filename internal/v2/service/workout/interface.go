package workout

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (id int64, err error)
}
