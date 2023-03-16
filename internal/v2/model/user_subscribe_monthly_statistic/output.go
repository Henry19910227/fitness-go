package user_subscribe_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_subscribe_monthly_statistics"
}
