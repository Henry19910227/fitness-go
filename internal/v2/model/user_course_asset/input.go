package user_course_asset

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/required"
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
	optional.UserIDField
	optional.CourseIDField
}

type ListInput struct {
	optional.UserIDField
	optional.CourseIDField
	optional.AvailableField
	optional.SourceField
	PagingInput
	JoinInput
	WhereInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type DeleteInput struct {
	required.IDField
}
