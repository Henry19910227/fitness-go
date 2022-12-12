package api_create_diet

import (
	dietOptional "github.com/Henry19910227/fitness-go/internal/v2/field/diet/optional"
	rdaOptional "github.com/Henry19910227/fitness-go/internal/v2/field/rda/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	dietOptional.IDField
	dietOptional.ScheduleAtField
	dietOptional.CreateAtField
	dietOptional.UpdateAtField
	RDA *struct {
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
	} `json:"rda,omitempty"`
}
