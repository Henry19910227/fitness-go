package api_get_cms_banners

import (
	orderByRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/required"
	pageRequired "github.com/Henry19910227/fitness-go/internal/v2/field/paging/required"
)

// Input /v2/cms/banners [GET]
type Input struct {
	Query Query
}
type Query struct {
	pageRequired.PageField
	pageRequired.SizeField
	orderByRequired.OrderFieldField
	orderByRequired.OrderTypeField
}
