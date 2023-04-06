package course_create_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_create_monthly_statistics"
}
