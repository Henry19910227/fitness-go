package receipt

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	receiptOptional "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/optional"
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
	receiptOptional.IDField
	receiptOptional.OrderIDField
	receiptOptional.OriginalTransactionIDField
	receiptOptional.TransactionIDField
	PreloadInput
	JoinInput
	WhereInput
}

type ListInput struct {
	receiptOptional.OrderIDField
	receiptOptional.PaymentTypeField
	receiptOptional.TransactionIDField
	receiptOptional.OriginalTransactionIDField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}
