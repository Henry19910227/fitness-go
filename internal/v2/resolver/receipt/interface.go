package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/receipt/api_get_cms_order_receipts"
)

type Resolver interface {
	APIGetCMSOrderReceipts(input *api_get_cms_order_receipts.Input) (output api_get_cms_order_receipts.Output)
}
