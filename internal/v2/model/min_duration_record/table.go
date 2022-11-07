package min_duration_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/min_duration_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.DurationField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "min_duration_records"
}
