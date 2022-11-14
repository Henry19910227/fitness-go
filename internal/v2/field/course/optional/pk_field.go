package optional

type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" uri:"course_id" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
