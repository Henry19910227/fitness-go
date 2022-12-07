package rda

import (
	rdaOptional "github.com/Henry19910227/fitness-go/internal/v2/field/rda/optional"
)

type Table struct {
	rdaOptional.IDField
	rdaOptional.UserIDField
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
	rdaOptional.CreateAtField
	rdaOptional.UpdateAtField
}

func (Table) TableName() string {
	return "rdas"
}
