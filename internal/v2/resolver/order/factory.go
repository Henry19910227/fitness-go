package order

import (
	orderService "github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	orderSvc := orderService.NewService(db)
	return New(orderSvc)
}