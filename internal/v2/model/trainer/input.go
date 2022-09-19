package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/required"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	optional.UserIDField
	PreloadInput
}

type ListInput struct {
	optional.UserIDField
	PreloadInput
	PagingInput
	OrderByInput
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
