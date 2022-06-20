package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount int
	CourseID   []*base.GenerateSetting
}

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
