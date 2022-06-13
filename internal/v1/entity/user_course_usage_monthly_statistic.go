package entity

type UserCourseUsageMonthlyStatistic struct {
	ID                  int64  `gorm:"column:id"`                    // 報表id
	UserID              int64  `gorm:"column:user_id"`               // 教練id
	FreeUsageCount      int    `gorm:"column:free_usage_count"`      // 免費課表使用人數
	SubscribeUsageCount int    `gorm:"column:subscribe_usage_count"` // 訂閱課表使用人數
	ChargeUsageCount    int    `gorm:"column:charge_usage_count"`    // 付費課表使用人數
	Year                int    `gorm:"column:year"`                  // 年份
	Month               int    `gorm:"column:month"`                 // 月份
	CreateAt            string `gorm:"column:create_at"`             // 創建日期
	UpdateAt            string `gorm:"column:update_at"`             // 更新日期
}

func (UserCourseUsageMonthlyStatistic) TableName() string {
	return "user_course_usage_monthly_statistics"
}
