package api_get_cms_order_receipts

import (
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	receiptOptional "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/order/{order_id}/receipts [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	receiptOptional.IDField
	receiptOptional.PaymentTypeField
	receiptOptional.OriginalTransactionIDField
	receiptOptional.TransactionIDField
	receiptOptional.QuantityField
	receiptOptional.CreateAtField
	ProductLabel *struct {
		productLabelOptional.IDField
		productLabelOptional.NameField
		productLabelOptional.ProductIDField
		productLabelOptional.TwdField
	} `json:"product_label,omitempty"`
}
