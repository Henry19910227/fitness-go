package user_plan_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_plan_statistics"
}
