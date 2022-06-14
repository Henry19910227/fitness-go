package plan

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	CourseIDField
	PagingInput
	OrderByInput
	PreloadInput
}

type APIGetCMSPlansInput struct {
	CourseIDField
	PagingInput
	OrderByInput
}
