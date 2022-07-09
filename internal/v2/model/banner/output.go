package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
}
func (Output) TableName() string {
	return "banners"
}

// APICreateCMSBannerOutput /v2/cms/banner [POST]
type APICreateCMSBannerOutput struct {
	base.Output
	Data *APICreateCMSBannerData `json:"data,omitempty"`
}
type APICreateCMSBannerData struct {
	IDField
	CourseIDField
	UserIDField
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
type APIGetCMSBannersData struct {
	IDField
	CourseIDField
	UserIDField
	ImageField
	TypeField
	CreateAtField
	UpdateAtField
}
