package card

type Output struct {
	Table
}

func (Output) TableName() string {
	return "cards"
}
