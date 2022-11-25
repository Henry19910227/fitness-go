package api_upload_google_subscribe_receipts

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/google_subscribe_receipts [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	ReceiptItems []struct {
		ProductID   string `json:"product_id" binding:"required" example:"com.fitness.xxx"`       // 產品id
		ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
	} `json:"receipt_items"`
}
