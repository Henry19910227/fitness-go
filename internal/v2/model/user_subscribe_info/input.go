package user_subscribe_info

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	UserIDOptional
}

type ListInput struct {
	UserIDOptional
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetUserSubscribeInfoInput /v2/user/subscribe_info [GET]
type APIGetUserSubscribeInfoInput struct {
	UserIDRequired
}
