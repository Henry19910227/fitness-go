package api_create_subscribe_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	subscribePlanOptional "github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_plan/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/subscribe_order [POST]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.IDField
	optional.UserIDField
	optional.QuantityField
	optional.OrderTypeField
	optional.OrderStatusField
	optional.CreateAtField
	optional.UpdateAtField
	OrderSubscribePlan *struct {
		SubscribePlan *struct {
			subscribePlanOptional.IDField
			subscribePlanOptional.PeriodField
			subscribePlanOptional.NameField
			ProductLabel *struct {
				productLabelOptional.IDField
				productLabelOptional.NameField
				productLabelOptional.ProductIDField
				productLabelOptional.TwdField
			} `json:"product_label,omitempty"`
		} `json:"subscribe_plan,omitempty"`
	} `json:"order_subscribe_plan,omitempty"`
}
