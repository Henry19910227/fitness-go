package receipt

import receiptOptional "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/optional"

type Table struct {
	receiptOptional.IDField
	receiptOptional.OrderIDField
	receiptOptional.PaymentTypeField
	receiptOptional.ReceiptTokenField
	receiptOptional.OriginalTransactionIDField
	receiptOptional.TransactionIDField
	receiptOptional.ProductIDField
	receiptOptional.QuantityField
	receiptOptional.CreateAtField
}

func (Table) TableName() string {
	return "receipts"
}
