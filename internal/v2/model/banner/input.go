package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
}

type DeleteInput struct {
	IDOptional
}

type ListInput struct {
	PagingInput
	OrderByInput
	PreloadInput
}

// APIGetBannersInput /v2/banners [GET]
type APIGetBannersInput struct {
	Query APIGetBannersQuery
}
type APIGetBannersQuery struct {
	PagingInput
}

// APICreateCMSBannerInput /v2/cms/banner [POST]
type APICreateCMSBannerInput struct {
	ImageFile *file.Input
	Form      APICreateCMSBannerForm
}
type APICreateCMSBannerForm struct {
	CourseIDOptional
	UserIDOptional
	TypeRequired
}

// APIDeleteCMSBannerInput /v2/cms/banner/{banner_id} [DELETE]
type APIDeleteCMSBannerInput struct {
	Uri APIDeleteCMSBannerUri
}
type APIDeleteCMSBannerUri struct {
	IDRequired
}

// APIGetCMSBannersInput /v2/cms/banners [GET]
type APIGetCMSBannersInput struct {
	Form APIGetCMSBannersForm
}
type APIGetCMSBannersForm struct {
	PagingInput
	OrderByInput
}
