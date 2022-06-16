package food_category

type Table struct {
	IDField
	TagField
	TitleField
	IsDeletedField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "food_categories"
}
