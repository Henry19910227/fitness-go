package user_subscribe_info

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userSubscribeInfoOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_info/optional"
	userSubscribeInfoRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_info/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/group"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/select_field"
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
type SelectInput = select_field.Input
type GroupInput = group.Input
type CustomOrderByInput = orderBy.CustomInput

type FindInput struct {
	userSubscribeInfoOptional.UserIDField
}

type ListInput struct {
	userSubscribeInfoOptional.UserIDField
	userSubscribeInfoOptional.StatusField
	userSubscribeInfoOptional.OrderIDField
	userSubscribeInfoOptional.OriginalTransactionIDField
	userSubscribeInfoOptional.PaymentTypeField
	SelectInput
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	GroupInput
	CustomOrderByInput
}

// APIGetUserSubscribeInfoInput /v2/user/subscribe_info [GET]
type APIGetUserSubscribeInfoInput struct {
	userSubscribeInfoRequired.UserIDField
}
