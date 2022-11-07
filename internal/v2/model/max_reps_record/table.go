package max_reps_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/max_reps_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.RepsField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "max_reps_records"
}
