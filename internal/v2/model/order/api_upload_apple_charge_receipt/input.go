package api_upload_apple_charge_receipt

import (
	orderRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/apple_charge_receipt [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}

type Body struct {
	orderRequired.OrderIDField
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}
