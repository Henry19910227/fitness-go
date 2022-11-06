package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set_log/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	optional.WorkoutLogIDField
	PagingInput
	OrderByInput
	PreloadInput
}
