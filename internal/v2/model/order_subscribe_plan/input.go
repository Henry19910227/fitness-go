package order_subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/order_subscribe_plan/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	optional.OrderIDField
}

type ListInput struct {
	optional.OrderIDField
	OrderByInput
	PagingInput
	PreloadInput
}
