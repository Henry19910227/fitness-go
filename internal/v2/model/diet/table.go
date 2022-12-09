package diet

import dietOptional "github.com/Henry19910227/fitness-go/internal/v2/field/diet/optional"

type Table struct {
	dietOptional.IDField
	dietOptional.UserIDField
	dietOptional.RdaIDField
	dietOptional.ScheduleAtField
	dietOptional.CreateAtField
	dietOptional.UpdateAtField
}

func (Table) TableName() string {
	return "diets"
}
