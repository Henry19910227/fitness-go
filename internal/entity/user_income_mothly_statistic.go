package entity

type UserIncomeMonthlyStatistic struct {
	ID       int64  `gorm:"column:id"`        // 報表id
	UserID   int64  `gorm:"column:user_id"`   // 教練id
	Income   int    `gorm:"column:income"`    // 收益
	Year     int    `gorm:"column:year"`      // 年份
	Month    int    `gorm:"column:month"`     // 月份
	CreateAt string `gorm:"column:create_at"` // 創建日期
	UpdateAt string `gorm:"column:update_at"` // 更新日期
}

func (UserIncomeMonthlyStatistic) TableName() string {
	return "user_income_monthly_statistics"
}
