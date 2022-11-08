package favorite_trainer

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_trainer/optional"

type Table struct {
	optional.UserIDField
	optional.TrainerIDField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "favorite_trainers"
}
