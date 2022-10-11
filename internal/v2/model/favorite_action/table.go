package favorite_action

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_aciton/optional"

type Table struct {
	optional.UserIDField
	optional.ActionIDField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "favorite_actions"
}
