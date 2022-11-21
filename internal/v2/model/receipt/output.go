package receipt

import (
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	receiptOptional "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
)

type Output struct {
	Table
	ProductLabel *product_label.Output `json:"product_label,omitempty" gorm:"foreignKey:product_id;references:product_id"`
}

func (Output) TableName() string {
	return "receipts"
}

// APIGetCMSOrderReceiptsOutput /v2/cms/order/{order_id}/receipts [GET]
type APIGetCMSOrderReceiptsOutput struct {
	base.Output
	Data   APIGetCMSOrderReceiptsData `json:"data"`
	Paging *paging.Output             `json:"paging,omitempty"`
}
type APIGetCMSOrderReceiptsData []*struct {
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
