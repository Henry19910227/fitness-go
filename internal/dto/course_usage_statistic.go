package dto

type CourseUsageStatistic struct {
	TotalFinishWorkoutCount int     `json:"total_finish_workout_count" example:"5500"`                                 // 歷史訓練總量(完成一次該課表的訓練，重複計算)
	UserFinishCount         int     `json:"user_finish_count" example:"3000"`                                          // 用戶使用人數(完成一次該課表的訓練，不重複計算)
	FinishCountAvg          int     `json:"finish_count_avg" example:"25"`                                             // 平均完成訓練次數(同一訓練不重複)
	MaleFinishCount         int     `json:"male_finish_count" gorm:"column:male_finish_count" example:"1000"`          // 男生使用人數
	FemaleFinishCount       int     `json:"female_finish_count" gorm:"column:female_finish_count" example:"1000"`      // 女生使用人數
	Age13to17Count          int     `json:"age_13_17_count" gorm:"column:age_13_17_count" example:"500"`               // 13-17歲使用人數
	Age18to24Count          int     `json:"age_18_24_count" gorm:"column:age_18_24_count" example:"100"`               // 18-24歲使用人數
	Age25to34Count          int     `json:"age_25_34_count" gorm:"column:age_25_34_count" example:"300"`               // 25-34歲使用人數
	Age35to44Count          int     `json:"age_35_44_count" gorm:"column:age_35_44_count" example:"200"`               // 35-44歲使用人數
	Age45to54Count          int     `json:"age_45_54_count" gorm:"column:age_45_54_count" example:"600"`               // 45-54歲使用人數
	Age55to64Count          int     `json:"age_55_64_count" gorm:"column:age_55_64_count" example:"700"`               // 55-64歲使用人數
	Age65upCount            int     `json:"age_65_up_count" gorm:"column:age_65_up_count" example:"550"`               // 65+歲使用人數
	CreateAt                *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2021-05-28 11:00:00"` // 創建日期
	UpdateAt                *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2021-05-28 11:00:00"` // 更新日期
}

type CourseUsageStatisticSummary struct {
	TotalFinishWorkoutCount int `json:"total_finish_workout_count" example:"5500"` // 歷史訓練總量(完成一次該課表的訓練，重複計算)
	UserFinishCount         int `json:"user_finish_count" example:"1000"`          // 用戶使用人數(完成一次該課表的訓練，不重複計算)
	FinishCountAvg          int `json:"finish_count_avg" example:"25"`             // 平均完成訓練次數(同一訓練不重複)
}
