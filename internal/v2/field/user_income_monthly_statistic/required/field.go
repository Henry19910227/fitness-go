package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 統計 id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type IncomeField struct {
	Income int `json:"income" gorm:"column:income" binding:"required" example:"12000"` // 收益
}
type YearField struct {
	Year int `json:"year" gorm:"column:year" binding:"required" example:"2022"` // 年份
}
type MonthField struct {
	Month int `json:"month" gorm:"column:month" binding:"required" example:"5"` // 月份
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
