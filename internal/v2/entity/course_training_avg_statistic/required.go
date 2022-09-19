package course_training_avg_statistic

type CourseIDRequired struct {
	CourseID int64 `json:"course_id" binding:"required" example:"10"` //課表id
}

type RateRequired struct {
	Rate int `json:"rate" binding:"required" example:"100"` //平均訓練率
}
