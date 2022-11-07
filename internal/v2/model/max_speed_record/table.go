package max_speed_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/max_speed_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.SpeedField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "max_speed_records"
}
