package certificate

type Output struct {
	Table
}

func (Output) TableName() string {
	return "certificates"
}
