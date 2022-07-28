package banner

import (
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
	IDField
	ImageField
	TypeField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.NicknameField
	} `json:"trainer,omitempty"`
	Course *struct {
		course.IDField
		course.NameField
	} `json:"course,omitempty"`
}

// APICreateCMSBannerOutput /v2/cms/banner [POST]
type APICreateCMSBannerOutput struct {
	base.Output
	Data *APICreateCMSBannerData `json:"data,omitempty"`
}
type APICreateCMSBannerData struct {
	IDField
	ImageField
	TypeField
	CreateAtField
	UpdateAtField
}

// APIDeleteCMSBannerOutput /v2/cms/banner [DELETE]
type APIDeleteCMSBannerOutput struct {
	base.Output
}

// APIGetCMSBannersOutput /v2/cms/banners [GET]
type APIGetCMSBannersOutput struct {
	base.Output
	Data   APIGetCMSBannersData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSBannersData []*struct {
	IDField
	ImageField
	TypeField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.NicknameField
	} `json:"trainer,omitempty"`
	Course *struct {
		course.IDField
		course.NameField
	} `json:"course,omitempty"`
}
