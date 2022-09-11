package max_speed_record

type Output struct {
	Table
}

func (Output) TableName() string {
	return "max_speed_records"
}
