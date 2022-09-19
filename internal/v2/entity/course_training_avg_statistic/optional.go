package course_training_avg_statistic

type CourseIDOptional struct {
	CourseID *int64 `json:"course_id,omitempty" form:"course_id" binding:"omitempty" example:"10"` //課表id
}
