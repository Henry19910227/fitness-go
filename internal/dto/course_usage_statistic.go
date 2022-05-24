package dto

type CourseUsageStatisticSummary struct {
	TotalFinishWorkoutCount int   `json:"total_finish_workout_count" example:"5500"` // 歷史訓練總量(完成一次該課表的訓練，重複計算)
	UserFinishCount         int   `json:"user_finish_count" example:"1000"`                   // 用戶使用人數(完成一次該課表的訓練，不重複計算)
	FinishCountAvg          int   `json:"finish_count_avg" example:"25"`                       // 平均完成訓練次數(同一訓練不重複)
}
