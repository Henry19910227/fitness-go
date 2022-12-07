package rda

type Output struct {
	Table
}

func (Output) TableName() string {
	return "rdas"
}
