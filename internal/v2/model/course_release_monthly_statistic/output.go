package course_release_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_release_monthly_statistics"
}
