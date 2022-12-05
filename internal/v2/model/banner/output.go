package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/banner/optional"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
)

type Output struct {
	Table
	Trainer *trainer.Output `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:user_id"` // 教練
	Course  *course.Output  `json:"course,omitempty" gorm:"foreignKey:id;references:course_id"`     // 課表
}

func (Output) TableName() string {
	return "banners"
}

// APIGetBannersOutput /v2/banners [GET]
type APIGetBannersOutput struct {
	base.Output
	Data   APIGetBannersData `json:"data"`
	Paging *paging.Output    `json:"paging,omitempty"`
}
type APIGetBannersData []*struct {
	optional.IDField
	optional.UrlField
	optional.ImageField
	optional.TypeField
	optional.CreateAtField
	optional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.NicknameField
	} `json:"trainer,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
}

// APIDeleteCMSBannerOutput /v2/cms/banner [DELETE]
type APIDeleteCMSBannerOutput struct {
	base.Output
}