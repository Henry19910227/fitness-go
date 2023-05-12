package subscribe_log

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/group"
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
type GroupInput = group.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type FindInput struct {
	optional.IDField
	optional.OriginalTransactionIDField
	optional.TransactionIDField
	optional.TypeField
	PreloadInput
}

type ListInput struct {
	optional.OriginalTransactionIDField
	optional.TransactionIDField
	optional.TypeField
	JoinInput
	GroupInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}
