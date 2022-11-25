package api_upload_google_charge_receipt

import orderRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order/required"

type Input struct {
	orderRequired.UserIDField
	Body Body
}

type Body struct {
	orderRequired.OrderIDField
	ProductID   string `json:"product_id" binding:"required" example:"com.fitness.xxx"`       // 產品id
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}
