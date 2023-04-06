package course_create_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course_create_monthly_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course_create_monthly_statistic/required"
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
	optional.YearField
	optional.MonthField
	PreloadInput
}

type ListInput struct {
	optional.YearField
	optional.MonthField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type StatisticInput struct {
	required.YearField
	required.MonthField
}
