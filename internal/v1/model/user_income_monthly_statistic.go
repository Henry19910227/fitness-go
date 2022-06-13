package model

type UserIncomeMonthlyStatisticResult struct {
	UserID int64 `gorm:"column:user_id"`
	Year   int   `gorm:"column:year"`
	Month  int64 `gorm:"column:month"`
	Income int   `gorm:"column:income"`
}
