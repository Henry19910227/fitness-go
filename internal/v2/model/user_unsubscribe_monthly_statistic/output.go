package user_unsubscribe_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_unsubscribe_monthly_statistics"
}
