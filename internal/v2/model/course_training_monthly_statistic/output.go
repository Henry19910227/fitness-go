package course_training_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_training_monthly_statistics"
}
