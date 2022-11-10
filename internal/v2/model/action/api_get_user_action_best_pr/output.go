package api_get_user_action_best_pr

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	maxDistanceRequired "github.com/Henry19910227/fitness-go/internal/v2/field/max_distance_record/required"
	maxRepsRequired "github.com/Henry19910227/fitness-go/internal/v2/field/max_reps_record/required"
	maxRMRequired "github.com/Henry19910227/fitness-go/internal/v2/field/max_rm_record/required"
	maxSpeedRequired "github.com/Henry19910227/fitness-go/internal/v2/field/max_speed_record/required"
	maxWeightRequired "github.com/Henry19910227/fitness-go/internal/v2/field/max_weight_record/required"
	minDurationRequired "github.com/Henry19910227/fitness-go/internal/v2/field/min_duration_record/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/user/action/{action_id}/best_personal_record [GET] 獲取用戶個人動作最佳紀錄 API
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.IDField
	MaxDistanceRecord struct {
		maxDistanceRequired.DistanceField
		maxDistanceRequired.UpdateAtField
	} `json:"max_distance_record"`
	MaxRepsRecord struct {
		maxRepsRequired.RepsField
		maxRepsRequired.UpdateAtField
	} `json:"max_reps_record"`
	MaxRMRecord struct{
		maxRMRequired.RMField
		maxRMRequired.UpdateAtField
	} `json:"max_rm_record"`
	MaxSpeedRecord struct{
		maxSpeedRequired.SpeedField
		maxSpeedRequired.UpdateAtField
	} `json:"max_speed_record"`
	MaxWeightRecord struct{
		maxWeightRequired.WeightField
		maxWeightRequired.UpdateAtField
	} `json:"max_weight_record"`
	MinDurationRecord struct{
		minDurationRequired.DurationField
		minDurationRequired.UpdateAtField
	} `json:"min_duration_record"`
}
