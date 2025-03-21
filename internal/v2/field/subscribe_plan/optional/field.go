package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` //訂閱項目id
}
type ProductLabelIDField struct {
	ProductLabelID *int64 `json:"product_label_id,omitempty" gorm:"column:product_label_id" binding:"omitempty" example:"1"` //產品標籤id
}
type PeriodField struct {
	Period *int `json:"period,omitempty" gorm:"column:period" binding:"omitempty" example:"12"` //週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" binding:"omitempty" example:"金牌課表"` //銷售名稱
}
type EnableField struct {
	Enable *int `json:"enable,omitempty" gorm:"column:enable" binding:"omitempty" example:"1"` //是否啟用
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
