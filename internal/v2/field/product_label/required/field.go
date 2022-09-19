package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 產品標籤id
}
type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required" example:"金卡會員(月)"` // 產品名稱
}
type ProductIDField struct {
	ProductID string `json:"product_id" gorm:"column:product_id" binding:"required" example:"com.fitness.gold_member_month"` // 產品id
}
type TwdField struct {
	Twd int `json:"twd" gorm:"column:twd" binding:"required" example:"500"` // 台幣價格
}
