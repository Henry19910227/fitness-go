package trainer_status_update_log

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainer_status_update_logs"
}
