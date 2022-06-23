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
}

type APIUpdateCMSTrainerAvatarInput struct {
	UserIDRequired
	CoverNamed string
	File       multipart.File
}