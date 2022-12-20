package user_course_usage_monthly_statistic

type Output struct {
	Table
}
func (Output) TableName() string {
	return "user_course_usage_monthly_statistics"
}
