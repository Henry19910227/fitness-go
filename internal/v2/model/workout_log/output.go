package workout_log

type Output struct {
	Table
}

func (Output) TableName() string {
	return "workout_logs"
}
