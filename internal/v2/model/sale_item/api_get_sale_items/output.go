package api_get_sale_items

import (
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/sale_items [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	optional.IDField
	optional.NameField
	ProductLabel *struct {
		productLabelOptional.IDField
		productLabelOptional.ProductIDField
		productLabelOptional.TwdField
	} `json:"product_label,omitempty"`
}
