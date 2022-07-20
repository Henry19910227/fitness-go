package order_course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (err error)
}
