package trainer_album

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainer_albums"
}
