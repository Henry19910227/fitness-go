package min_duration_record

type Output struct {
	Table
}

func (Output) TableName() string {
	return "min_duration_records"
}
