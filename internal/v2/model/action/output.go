package action

type Table struct {
	IDField
	CourseIDField
	NameField
	SourceField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "actions"
}
