package card

import "github.com/Henry19910227/fitness-go/internal/v2/field/card/optional"


type Table struct {
	optional.UserIDField
	optional.CardIDField
	optional.FrontImageField
	optional.BackImageField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "cards"
}
