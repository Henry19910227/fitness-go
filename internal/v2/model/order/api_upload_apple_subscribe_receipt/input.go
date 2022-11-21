package api_upload_apple_subscribe_receipt

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/apple_subscribe_receipt [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}
