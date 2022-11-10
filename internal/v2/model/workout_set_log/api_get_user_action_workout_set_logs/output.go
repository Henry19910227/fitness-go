package api_get_user_action_workout_set_logs

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/user/action/{action_id}/workout_set_logs [GET] 以日期獲取動作訓練組紀錄 API
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	optional.IDField
	optional.WorkoutLogIDField
	optional.WorkoutSetIDField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
}
