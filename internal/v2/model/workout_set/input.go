package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// GenerateInput Test Input
type GenerateInput struct {
	DataAmount int
	WorkoutID  []*base.GenerateSetting
}

type ListInput struct {
	WorkoutIDField
	PagingInput
	PreloadInput
}

type APIGetCMSWorkoutSetsInput struct {
	WorkoutIDField
	PagingInput
}
