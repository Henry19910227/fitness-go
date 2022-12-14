package api_calculate_rda

import (
	rdaOptional "github.com/Henry19910227/fitness-go/internal/v2/field/rda/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/calculate_rda [POST]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	rdaOptional.TDEEField
	rdaOptional.CalorieField
	rdaOptional.ProteinField
	rdaOptional.FatField
	rdaOptional.CarbsField
	rdaOptional.GrainField
	rdaOptional.VegetableField
	rdaOptional.FruitField
	rdaOptional.MeatField
	rdaOptional.DairyField
	rdaOptional.NutField
}
