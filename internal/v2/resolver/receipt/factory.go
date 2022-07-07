package receipt

import (
	receiptService "github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	receiptSvc := receiptService.NewService(db)
	return New(receiptSvc)
}
