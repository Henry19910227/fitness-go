package api_get_cms_banners

import (
	bannerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/banner/optional"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/banners [GET]
type Output struct {
	base.Output
	Data   Data `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type Data []*struct {
	bannerOptional.IDField
	bannerOptional.ImageField
	bannerOptional.TypeField
	bannerOptional.CreateAtField
	bannerOptional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.NicknameField
	} `json:"trainer,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
}
