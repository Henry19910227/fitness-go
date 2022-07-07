package product_label

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` // 產品標籤id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" example:"金卡會員(月)"` // 產品名稱
}
type ProductIDField struct {
	ProductID *string `json:"product_id,omitempty" gorm:"column:product_id" example:"com.fitness.gold_member_month"` // 產品id
}
type TwdField struct {
	Twd *int `json:"twd,omitempty" gorm:"column:twd" example:"500"` // 台幣價格
}

type Table struct {
	IDField
	NameField
	ProductIDField
	TwdField
}

func (Table) TableName() string {
	return "product_labels"
}
