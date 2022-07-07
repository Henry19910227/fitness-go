package receipt

import (
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

// APIGetCMSReceiptsOutput /v2/cms/receipts [GET]
type APIGetCMSReceiptsOutput struct {
	base.Output
	Data   APIGetCMSReceiptsData `json:"data"`
	Paging *paging.Output        `json:"paging,omitempty"`
}
type APIGetCMSReceiptsData []*struct {
	IDField
	PaymentTypeField
	OriginalTransactionIDField
	TransactionIDField
	QuantityField
	CreateAtField
	ProductLabel *struct {
		product_label.IDField
		product_label.NameField
		product_label.ProductIDField
		product_label.TwdField
	} `json:"product_label,omitempty"`
}
