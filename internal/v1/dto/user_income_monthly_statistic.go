package dto

type UserIncomeMonthlyStatistic struct {
	Income   int    `json:"income" gorm:"column:income" example:"12000"`                               // 收益
	Year     int    `json:"year,omitempty" gorm:"column:year" example:"5"`                             // 年份
	Month    int    `json:"month,omitempty" gorm:"column:month" example:"25"`                          // 月份
	CreateAt string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-05-25 11:00:00"` // 創建日期
	UpdateAt string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-05-25 12:00:00"` // 更新日期
}

func (UserIncomeMonthlyStatistic) TableName() string {
	return "user_income_monthly_statistics"
}
