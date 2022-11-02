package card

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/card"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (err error)
}
