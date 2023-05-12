package subscribe_log

import "github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_log/optional"

type Table struct {
	optional.IDField
	optional.OriginalTransactionIDField
	optional.TransactionIDField
	optional.PurchaseDateField
	optional.ExpiresDateField
	optional.TypeField
	optional.MsgField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "subscribe_logs"
}
