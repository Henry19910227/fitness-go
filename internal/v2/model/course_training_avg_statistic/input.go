package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
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

type FindInput struct {
	optional.CourseIDField
	PreloadInput
}

type ListInput struct {
	optional.CourseIDField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}
