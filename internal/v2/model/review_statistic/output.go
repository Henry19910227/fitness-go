package review_statistic

type Output struct {
	Table
}
func (Output) TableName() string {
	return "review_statistics"
}
