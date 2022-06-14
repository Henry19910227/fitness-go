package sale_item

import "github.com/Henry19910227/fitness-go/internal/v2/model/product_label"

type Table struct {
	IDField
	ProductLabelIDField
	TypeField
	EnableField
	NameField
	CreateAtField
	UpdateAtField
	ProductLabel *product_label.Table `json:"product_label,omitempty" gorm:"foreignKey:id;references:product_label_id"` // 產品標籤
}

func (Table) TableName() string {
	return "sale_items"
}
