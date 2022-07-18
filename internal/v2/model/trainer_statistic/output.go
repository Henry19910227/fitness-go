package trainer_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainer_statistics"
}
