package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input
type WhereInput = where.Input
type JoinInput = join.Input
type CustomOrderByInput = orderBy.CustomInput

type FindInput struct {
	optional.UserIDField
	PreloadInput
}

type ListInput struct {
	optional.UserIDField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type FavoriteListInput struct {
	optional.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetTrainerProfileInput /v2/trainer/profile [PATCH]
type APIGetTrainerProfileInput struct {
	required.UserIDField
}

// APIGetStoreTrainerInput /v2/store/trainer/{user_id} [GET]
type APIGetStoreTrainerInput struct {
	Uri APIGetStoreTrainerUri
}
type APIGetStoreTrainerUri struct {
	required.UserIDField
}

// APIGetStoreTrainersInput /v2/store/trainers [GET]
type APIGetStoreTrainersInput struct {
	required.UserIDField
	Query APIGetStoreTrainersQuery
}
type APIGetStoreTrainersQuery struct {
	OrderField *string `json:"order_field" form:"order_field" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	PagingInput
}

// APIUpdateCMSTrainerAvatarInput /v2/cms/trainer/avatar [PATCH]
type APIUpdateCMSTrainerAvatarInput struct {
	required.UserIDField
	CoverNamed string
	File       multipart.File
}

// APIGetFavoriteTrainersInput /v2/favorite/trainers [GET]
type APIGetFavoriteTrainersInput struct {
	required.UserIDField
	Form APIGetFavoriteTrainersForm
}
type APIGetFavoriteTrainersForm struct {
	PagingInput
}
