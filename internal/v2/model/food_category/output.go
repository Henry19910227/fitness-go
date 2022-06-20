package food_category

type Output struct {
	Table
}

func (Output) TableName() string {
	return "food_categories"
}
