package max_distance_record

type Output struct {
	Table
}

func (Output) TableName() string {
	return "max_distance_records"
}