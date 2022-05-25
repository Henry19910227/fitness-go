package entity

type CourseUsageStatistic struct {
	ID                      int64  `gorm:"column:id"`                         // 報表id
	CourseID                int64  `gorm:"column:course_id"`                  // 課表id
	TotalFinishWorkoutCount int    `gorm:"column:total_finish_workout_count"` // 歷史訓練總量(完成一次該課表的訓練，重複計算)
	UserFinishCount         int    `gorm:"column:user_finish_count"`          // 用戶使用人數(完成一次該課表的訓練，不重複計算)
	MaleFinishCount         int    `gorm:"column:male_finish_count"`          // 男生使用人數
	FemaleFinishCount       int    `gorm:"column:female_finish_count"`        // 女生使用人數
	FinishCountAvg          int    `gorm:"column:finish_count_avg"`           // 平均完成訓練次數(同一訓練不重複)
	Age13to17Count          int    `gorm:"column:age_13_17_count"`            // 13-17歲使用人數
	Age18to24Count          int    `gorm:"column:age_18_24_count"`            // 18-24歲使用人數
	Age25to34Count          int    `gorm:"column:age_25_34_count"`            // 25-34歲使用人數
	Age35to44Count          int    `gorm:"column:age_35_44_count"`            // 35-44歲使用人數
	Age45to54Count          int    `gorm:"column:age_45_54_count"`            // 45-54歲使用人數
	Age55to64Count          int    `gorm:"column:age_55_64_count"`            // 55-64歲使用人數
	Age65upCount            int    `gorm:"column:age_65_up_count"`            // 65+歲使用人數
	CreateAt                string `gorm:"column:create_at"`                  // 創建日期
	UpdateAt                string `gorm:"column:update_at"`                  // 更新日期
}

func (CourseUsageStatistic) TableName() string {
	return "course_usage_statistics"
}
