package certificate

import "github.com/Henry19910227/fitness-go/internal/v2/field/certificate/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.NameField
	optional.ImageField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "certificates"
}
