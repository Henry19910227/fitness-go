package api_get_cms_trainers

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/trainer/{user_id} [PATCH]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	optional.UserIDField
	optional.NicknameField
	optional.TrainerStatusField
	optional.EmailField
	optional.YoutubeURLField
	optional.UpdateAtField
}
