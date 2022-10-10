package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	optional.UserIDField
	optional.CourseIDField
	optional.AvailableField
	PagingInput
	PreloadInput
	OrderByInput
}