package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/sale_item"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := sale_item.NewResolver(db)
	return New(resolver)
}
