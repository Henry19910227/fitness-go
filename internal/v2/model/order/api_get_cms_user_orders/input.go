package api_get_cms_user_orders

import (
	orderByOptional "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/cms/user/{user_id}/orders [GET]
type Input struct {
	userRequired.UserIDField
	Uri Uri
	Query Query
}
type Uri struct {
	userRequired.UserIDField
}
type Query struct {
	orderByOptional.OrderFieldField
	orderByOptional.OrderTypeField
	pagingOptional.PageField
	pagingOptional.SizeField
}
