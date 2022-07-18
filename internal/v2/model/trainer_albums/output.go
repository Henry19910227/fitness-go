package trainer_albums

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainer_albums"
}
