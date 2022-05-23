package model

type UserCourseUsageMonthlyStatisticResult struct {
	UserID int64 `gorm:"column:user_id"`
	Year int `gorm:"column:year"`
	Month int64 `gorm:"column:month"`
	Value int `gorm:"column:value"`
}
