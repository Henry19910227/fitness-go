package coursedto

type CreateResult struct {
	ID int64 `json:"course_id" example:"1"` //課表 id
}

type CreateCourseParam struct {
	Name string
	Level int
	Category int
	CategoryOther string
	ScheduleType int
}
