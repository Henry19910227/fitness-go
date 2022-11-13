package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	saleItemService := sale_item.NewService(db)
	return New(saleItemService)
}
