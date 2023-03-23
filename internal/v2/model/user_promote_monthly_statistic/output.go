package user_promote_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_promote_monthly_statistics"
}
