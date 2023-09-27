package api_get_cms_order_receipts

import (
	orderByRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/required"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	receiptRequired "github.com/Henry19910227/fitness-go/internal/v2/field/receipt/required"
)

// Input /v2/cms/order/{order_id}/receipts [GET]
type Input struct {
	Uri   Uri
	Query Query
}
type Uri struct {
	receiptRequired.OrderIDField
}
type Query struct {
	pagingOptional.PageField
	pagingOptional.SizeField
	orderByRequired.OrderFieldField
	orderByRequired.OrderTypeField
}
