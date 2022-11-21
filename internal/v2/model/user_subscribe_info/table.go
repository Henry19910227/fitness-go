package user_subscribe_info

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_info/optional"

type Table struct {
	optional.UserIDField
	optional.OrderIDField
	optional.OriginalTransactionIDField
	optional.StatusField
	optional.PaymentTypeField
	optional.StartDateField
	optional.ExpiresDateField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "user_subscribe_infos"
}
