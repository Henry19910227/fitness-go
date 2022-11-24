package api_upload_apple_subscribe_receipts

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/apple_subscribe_receipts [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	ReceiptDatas []string `json:"receipt_datas" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}
