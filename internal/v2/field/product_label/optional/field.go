package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 產品標籤id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" binding:"omitempty" example:"金卡會員(月)"` // 產品名稱
}
type ProductIDField struct {
	ProductID *string `json:"product_id,omitempty" gorm:"column:product_id" binding:"omitempty" example:"com.fitness.gold_member_month"` // 產品id
}
type TwdField struct {
	Twd *int `json:"twd,omitempty" gorm:"column:twd" binding:"omitempty" example:"500"` // 台幣價格
}
