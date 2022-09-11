package max_rm_record

type Output struct {
	Table
}

func (Output) TableName() string {
	return "max_rm_records"
}