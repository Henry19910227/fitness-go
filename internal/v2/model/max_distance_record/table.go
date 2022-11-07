package max_distance_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/max_distance_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.DistanceField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "max_distance_records"
}
