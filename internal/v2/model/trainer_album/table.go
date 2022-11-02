package trainer_album

import "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_album/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.PhotoField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "trainer_albums"
}
