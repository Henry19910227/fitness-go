package subscribe_log

type Output struct {
	Table
}

func (Output) TableName() string {
	return "subscribe_logs"
}
