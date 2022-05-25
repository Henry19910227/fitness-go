package dto

type UserCourseUsageMonthlyStatistic struct {
	FreeUsageCount      int    `json:"free_usage_count" gorm:"column:free_usage_count" example:"400"`      // 免費課表使用人數
	SubscribeUsageCount int    `json:"subscribe_usage_count" gorm:"column:subscribe_usage_count" example:"60"` // 訂閱課表使用人數
	ChargeUsageCount    int    `json:"charge_usage_count" gorm:"column:charge_usage_count" example:"50"`    // 付費課表使用人數
	Year                *int    `json:"year,omitempty" gorm:"column:year" example:"2022"`                  // 年份
	Month               *int    `json:"month,omitempty" gorm:"column:month" example:"5"`                 // 月份
	CreateAt            *string `json:"create_at,omitempty"gorm:"column:create_at" example:"2022-05-20 12:00:00"`  // 創建日期
	UpdateAt            *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-05-21 12:00:00"` // 更新日期
}

func (UserCourseUsageMonthlyStatistic) TableName() string {
	return "user_course_usage_monthly_statistics"
}
