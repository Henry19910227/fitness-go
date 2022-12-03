package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/service/banner_order"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bannerOrderService := banner_order.NewService(db)
	bannerService := banner.NewService(db)
	return New(bannerOrderService, bannerService)
}
