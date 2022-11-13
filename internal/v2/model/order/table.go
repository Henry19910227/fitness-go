package order

import "github.com/Henry19910227/fitness-go/internal/v2/field/order/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.QuantityField
	optional.OrderTypeField
	optional.OrderStatusField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "orders"
}
