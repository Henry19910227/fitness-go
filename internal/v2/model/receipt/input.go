package receipt

import (
	receiptOptional "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/optional"
	receiptRequired "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/required"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	receiptOptional.IDField
}

type ListInput struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` // 用戶 id
	receiptOptional.OrderIDField
	receiptOptional.PaymentTypeField
	receiptOptional.TransactionIDField
	receiptOptional.OriginalTransactionIDField
	HaveReceiptToken *int `json:"have_receipt_token,omitempty" binding:"omitempty" example:"1"` // 是否擁有 ReceiptToken(0:否/1:是)
	PreloadInput
	OrderByInput
	PagingInput
}

// APIGetCMSOrderReceiptsInput /v2/cms/order/{order_id}/receipts [GET]
type APIGetCMSOrderReceiptsInput struct {
	Uri  APIGetCMSOrderReceiptsUri
	Form APIGetCMSOrderReceiptsForm
}
type APIGetCMSOrderReceiptsUri struct {
	receiptRequired.OrderIDField
}
type APIGetCMSOrderReceiptsForm struct {
	OrderByInput
	PagingInput
}
