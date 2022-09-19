package product_label

import "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"

type Table struct {
	optional.IDField
	optional.NameField
	optional.ProductIDField
	optional.TwdField
}

func (Table) TableName() string {
	return "product_labels"
}
