package course_category_training_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_category_training_monthly_statistics"
}
