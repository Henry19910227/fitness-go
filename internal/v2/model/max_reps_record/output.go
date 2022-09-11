package max_reps_record

type Output struct {
	Table
}

func (Output) TableName() string {
	return "max_reps_records"
}
