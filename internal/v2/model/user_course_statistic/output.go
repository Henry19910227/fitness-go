package user_course_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_course_statistics"
}
