package sale_item

import "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"

type Table struct {
	optional.IDField
	optional.ProductLabelIDField
	optional.TypeField
	optional.EnableField
	optional.NameField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "sale_items"
}
