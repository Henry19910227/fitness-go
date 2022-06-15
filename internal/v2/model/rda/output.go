package rda

type Table struct {
	IDField
	UserIDField
	TDEEField
	CalorieField
	ProteinField
	FatField
	CarbsField
	GrainField
	VegetableField
	FruitField
	MeatField
	DairyField
	NutField
	CreateAtField
}

func (Table) TableName() string {
	return "rdas"
}
