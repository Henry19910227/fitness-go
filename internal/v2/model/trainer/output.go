package trainer

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainers"
}

// APIUpdateCMSTrainerAvatarOutput /v2/cms/trainer/{user_id}/avatar [PATCH] 獲取課表詳細 API
type APIUpdateCMSTrainerAvatarOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}
