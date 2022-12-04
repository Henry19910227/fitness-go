package api_get_cms_banners

import (
	orderByRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/required"
	pageOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
)

// Input /v2/cms/banners [GET]
type Input struct {
	Query Query
}
type Query struct {
	pageOptional.PageField
	pageOptional.SizeField
	orderByRequired.OrderFieldField
	orderByRequired.OrderTypeField
}
