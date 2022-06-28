package body_image

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

type ListInput struct {
	UserIDOptional
	PagingInput
	OrderByInput
}

// APIGetBodyImagesInput /body_images [GET] 獲取體態照片列表
type APIGetBodyImagesInput struct {
	UserIDRequired
	Query APIGetBodyImagesQuery
}
type APIGetBodyImagesQuery struct {
	PagingInput
}

// APICreateBodyImageInput /body_image [POST] 新增體態照片
type APICreateBodyImageInput struct {
	UserIDRequired
	ImageFile *file.Input
}
