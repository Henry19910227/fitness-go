package api_get_user_action_workout_set_logs

import (
	actionRequired "github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type PagingInput = paging.Input

// Input /v2/user/action/{action_id}/workout_set_logs [GET] 以日期獲取動作訓練組紀錄 API
type Input struct {
	userRequired.UserIDField
	Uri   Uri
	Query Query
}
type Uri struct {
	actionRequired.ActionIDField
}
type Query struct {
	StartDate string `form:"start_date" binding:"required,datetime=2006-01-02" example:"區間開始日期"` //區間開始日期
	EndDate   string `form:"end_date" binding:"required,datetime=2006-01-02" example:"區間結束日期"`   //區間結束日期
	PagingInput
}
