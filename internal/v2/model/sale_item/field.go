package sale_item

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` // 銷售id
}
type ProductLabelIDField struct {
	ProductLabelID *int `json:"product_label_id,omitempty" gorm:"column:product_label_id" example:"2"` // 產品標籤id
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type" example:"3"` // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
}
type EnableField struct {
	Enable *int `json:"enable,omitempty" gorm:"column:enable" example:"1"` // 是否啟用
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" example:"銅級課表 "` // 銷售名稱
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-12 00:00:00"` // 更新時間
}
