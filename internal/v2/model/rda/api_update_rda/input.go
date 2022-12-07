package api_update_rda

import (
	rdaRequired "github.com/Henry19910227/fitness-go/internal/v2/field/rda/required"
)

// Input /v2/rda [PUT]
type Input struct {
	rdaRequired.UserIDField
	Body Body
}
type Body struct {
	rdaRequired.TDEEField
	rdaRequired.CalorieField
	rdaRequired.ProteinField
	rdaRequired.FatField
	rdaRequired.CarbsField
	rdaRequired.GrainField
	rdaRequired.VegetableField
	rdaRequired.FruitField
	rdaRequired.MeatField
	rdaRequired.DairyField
	rdaRequired.NutField
}
