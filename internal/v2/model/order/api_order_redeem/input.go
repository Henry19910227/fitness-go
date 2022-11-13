package api_order_redeem

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/order/{order_id}/redeem [POST]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.OrderIDField
}
