package course_status_update_log

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_status_update_logs"
}
