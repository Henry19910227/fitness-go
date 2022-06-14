package workout

type Table struct {
	IDField
	PlanIDField
	NameField
	EquipmentField
	StartAudioField
	EndAudioField
	WorkoutSetCountField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "workouts"
}
