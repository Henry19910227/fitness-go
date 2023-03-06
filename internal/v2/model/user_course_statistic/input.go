package user_course_statistic

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
)

type PagingInput = struct {
	pagingOptional.PageField
	pagingOptional.SizeField
}
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type ListInput struct {
	optional.IDField
	optional.UserIDField
	optional.CourseIDField
	optional.FinishWorkoutCountField
	optional.TotalFinishWorkoutCountField
	optional.DurationField
	optional.UpdateAtField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type FindInput struct {
	optional.IDField
	PreloadInput
}
