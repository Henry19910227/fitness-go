package api_update_cms_banner_orders

import "github.com/Henry19910227/fitness-go/internal/v2/field/banner_order/required"

// Input /v2/cms/banner_orders [PUT] 修改banner排序
type Input struct {
	Body Body
}
type Body struct {
	BannerOrders []*struct{
		required.BannerIDField
		required.SeqField
	} `json:"banner_orders"`
}
