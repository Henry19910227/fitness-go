package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 統計 id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type FreeUsageCountField struct {
	FreeUsageCount int `json:"free_usage_count" gorm:"column:free_usage_count" binding:"required" example:"400"`      // 免費課表使用人數
}
type SubscribeUsageCountField struct {
	SubscribeUsageCount int `json:"subscribe_usage_count" gorm:"column:subscribe_usage_count" binding:"required" example:"60"` // 訂閱課表使用人數
}
type ChargeUsageCountField struct {
	ChargeUsageCount int `json:"charge_usage_count" gorm:"column:charge_usage_count" binding:"required" example:"50"`    // 付費課表使用人數
}
type YearField struct {
	Year int    `json:"year" gorm:"column:year" binding:"required" example:"2022"`                  // 年份
}
type MonthField struct {
	Month int    `json:"month" gorm:"column:month" binding:"required" example:"5"`                 // 月份
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
