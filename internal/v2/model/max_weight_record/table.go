package max_weight_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/max_weight_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.WeightField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "max_weight_records"
}
