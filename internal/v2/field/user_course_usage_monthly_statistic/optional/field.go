package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 統計 id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type FreeUsageCountField struct {
	FreeUsageCount *int `json:"free_usage_count,omitempty" gorm:"column:free_usage_count" binding:"omitempty" example:"400"`      // 免費課表使用人數
}
type SubscribeUsageCountField struct {
	SubscribeUsageCount *int `json:"subscribe_usage_count,omitempty" gorm:"column:subscribe_usage_count" binding:"omitempty" example:"60"` // 訂閱課表使用人數
}
type ChargeUsageCountField struct {
	ChargeUsageCount *int `json:"charge_usage_count,omitempty" gorm:"column:charge_usage_count" binding:"omitempty" example:"50"`    // 付費課表使用人數
}
type YearField struct {
	Year *int    `json:"year,omitempty" gorm:"column:year" binding:"omitempty" example:"2022"`                  // 年份
}
type MonthField struct {
	Month *int    `json:"month,omitempty" gorm:"column:month" binding:"omitempty" example:"5"`                 // 月份
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}