package banner_order

import "github.com/Henry19910227/fitness-go/internal/v2/field/banner_order/optional"

type Table struct {
	optional.BannerIDField
	optional.SeqField
}

func (Table) TableName() string {
	return "banner_orders"
}