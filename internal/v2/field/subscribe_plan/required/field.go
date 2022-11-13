package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //訂閱項目id
}
type ProductLabelIDField struct {
	ProductLabelID int64 `json:"product_label_id" gorm:"column:product_label_id" binding:"required" example:"1"` //產品標籤id
}
type PeriodField struct {
	Period int `json:"period" gorm:"column:period" binding:"required" example:"12"` //週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
}
type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required" example:"金牌課表"` //銷售名稱
}
type EnableField struct {
	Enable int `json:"enable" gorm:"column:enable" binding:"required" example:"1"` //是否啟用
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
