package trainer_status_update_log

import "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_status_update_log/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.TrainerStatusField
	optional.CommentField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "trainer_status_update_logs"
}
