package course_usage_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_usage_statistics"
}
