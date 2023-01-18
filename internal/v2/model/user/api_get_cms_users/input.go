package api_get_cms_users

import (
	orderByOptional "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
)

// Input /v2/cms/users [GET]
type Input struct {
	Query Query
}
type Query struct {
	userOptional.UserIDField
	userOptional.NicknameField
	userOptional.EmailField
	userOptional.UserStatusField
	userOptional.UserTypeField
	orderByOptional.OrderFieldField
	orderByOptional.OrderTypeField
	pagingOptional.PageField
	pagingOptional.SizeField
}
