package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner_order/api_update_cms_banner_orders"
	"gorm.io/gorm"
)

type Resolver interface {
	APIUpdateCMSBannerOrders(tx *gorm.DB, input *api_update_cms_banner_orders.Input) (output api_update_cms_banner_orders.Output)
}
