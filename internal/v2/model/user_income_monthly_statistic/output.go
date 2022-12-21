package user_income_monthly_statistic

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_income_monthly_statistics"
}
