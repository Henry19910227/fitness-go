package max_rm_record

import "github.com/Henry19910227/fitness-go/internal/v2/field/max_rm_record/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.ActionIDField
	optional.RMField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "max_rm_records"
}
