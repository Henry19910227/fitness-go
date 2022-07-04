package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "trainers"
}

// APIGetFavoriteTrainersOutput /v2/favorite/trainers [GET] 獲取
type APIGetFavoriteTrainersOutput struct {
	base.Output
	Data   APIGetFavoriteTrainersData `json:"data"`
	Paging *paging.Output            `json:"paging,omitempty"`
}
type APIGetFavoriteTrainersData []*struct {
	UserIDField
	NicknameField
	SkillField
	AvatarField
}

// APIUpdateCMSTrainerAvatarOutput /v2/cms/trainer/{user_id}/avatar [PATCH] 獲取課表詳細 API
type APIUpdateCMSTrainerAvatarOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}
