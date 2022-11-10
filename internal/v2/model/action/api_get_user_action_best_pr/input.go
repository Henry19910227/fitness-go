package api_get_user_action_best_pr

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/user/action/{action_id}/best_personal_record [GET] 獲取用戶個人動作最佳紀錄 API
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.ActionIDField
}
