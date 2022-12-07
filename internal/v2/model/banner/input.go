package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/banner/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
)

type PagingInput = struct {
	pagingOptional.PageField
	pagingOptional.SizeField
}
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type FindInput struct {
	optional.IDField
}

type DeleteInput struct {
	optional.IDField
}

type ListInput struct {
	PagingInput
	JoinInput
	WhereInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

// APIGetBannersInput /v2/banners [GET]
type APIGetBannersInput struct {
	Query APIGetBannersQuery
}
type APIGetBannersQuery struct {
	PagingInput
}

// APIDeleteCMSBannerInput /v2/cms/banner/{banner_id} [DELETE]
type APIDeleteCMSBannerInput struct {
	Uri APIDeleteCMSBannerUri
}
type APIDeleteCMSBannerUri struct {
	required.IDField
}
