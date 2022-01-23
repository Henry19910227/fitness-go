package model

type ProductLabel struct {
	ID int64 `gorm:"column:id"` // 產品標籤id
	Name string `gorm:"column:name"` // 產品名稱
	ProductID string `gorm:"column:product_id"` // 產品id
	Twd int  `gorm:"column:twd"` // 台幣價格
}

func (ProductLabel) TableName() string {
	return "product_labels"
}
