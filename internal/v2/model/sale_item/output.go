package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
)

type Output struct {
	Table
	ProductLabel *product_label.Output `json:"product_label,omitempty" gorm:"foreignKey:id;references:product_label_id"` // 產品標籤
}

func (Output) TableName() string {
	return "sale_items"
}

func (o Output) ProductLabelOnSafe() product_label.Output {
	if o.ProductLabel != nil {
		return *o.ProductLabel
	}
	return product_label.Output{}
}
