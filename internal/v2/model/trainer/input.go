package trainer

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	UserIDOptional
	PreloadInput
}

type ListInput struct {
	UserIDOptional
	PreloadInput
	PagingInput
	OrderByInput
}

type FavoriteListInput struct {
	UserIDOptional
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetTrainerProfileInput /v2/trainer/profile [PATCH]
type APIGetTrainerProfileInput struct {
	UserIDRequired
}

// APIUpdateCMSTrainerAvatarInput /v2/cms/trainer/avatar [PATCH]
type APIUpdateCMSTrainerAvatarInput struct {
	UserIDRequired
	CoverNamed string
	File       multipart.File
}

// APIGetFavoriteTrainersInput /v2/favorite/trainers [GET]
type APIGetFavoriteTrainersInput struct {
	UserIDRequired
	Form APIGetFavoriteTrainersForm
}
type APIGetFavoriteTrainersForm struct {
	PagingInput
}
