package required

type CourseIDField struct {
	CourseID int64 `json:"course_id" uri:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
