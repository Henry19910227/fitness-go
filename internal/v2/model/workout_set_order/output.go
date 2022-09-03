package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "workout_set_orders"
}

// APIUpdateUserWorkoutSetOrdersOutput /v2/user/workout/{workout_id}/workout_set_orders [PUT] 修改訓練組的順序
type APIUpdateUserWorkoutSetOrdersOutput struct {
	base.Output
}