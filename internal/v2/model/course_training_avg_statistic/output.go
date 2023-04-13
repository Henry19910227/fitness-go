package course_training_avg_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_training_avg_statistics"
}
