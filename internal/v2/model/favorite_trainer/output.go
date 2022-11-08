package favorite_trainer

type Output struct {
	Table
}

func (Output) TableName() string {
	return "favorite_trainers"
}
