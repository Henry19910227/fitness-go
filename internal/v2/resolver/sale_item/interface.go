package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/sale_item/api_get_sale_items"
)

type Resolver interface {
	APIGetSaleItems(input *api_get_sale_items.Input) (output api_get_sale_items.Output)
}
