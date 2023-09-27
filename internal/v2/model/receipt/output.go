package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
)

type Output struct {
	Table
	ProductLabel *product_label.Output `json:"product_label,omitempty" gorm:"foreignKey:product_id;references:product_id"`
}

func (Output) TableName() string {
	return "receipts"
}
