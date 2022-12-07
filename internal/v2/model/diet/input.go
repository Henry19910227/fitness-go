package diet

import (
	dietOptional "github.com/Henry19910227/fitness-go/internal/v2/field/diet/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
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

type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
	RdaID      []*base.GenerateSetting
}

type FindInput struct {
	IDField
	preload.Input
}

type ListInput struct {
	dietOptional.UserIDField
	PagingInput
	JoinInput
	WhereInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}
