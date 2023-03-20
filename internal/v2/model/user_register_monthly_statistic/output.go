package user_register_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_register_monthly_statistics"
}
