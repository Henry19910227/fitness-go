package api_get_subscribe_plans

import (
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_plan/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/subscribe_plans [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data []*struct {
	optional.IDField
	optional.PeriodField
	optional.NameField
	ProductLabel *struct {
		productLabelOptional.IDField
		productLabelOptional.NameField
		productLabelOptional.ProductIDField
		productLabelOptional.TwdField
	} `json:"product_label,omitempty"`
}
