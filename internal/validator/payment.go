package validator

type CreateCourseOrderBody struct {
	CourseID int64 `json:"course_id" binding:"required" example:"1"` // 課表 id
}

type CreateSubscribeOrderBody struct {
	SubscribePlanID int64 `json:"subscribe_plan_id" binding:"required" example:"1"` // 訂閱方案 id
}

type VerifyReceiptBody struct {
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
	OrderID     string `json:"order_id" binding:"required" example:"202105201300687423"`      // 訂單 id
}

type VerifyGoogleReceiptBody struct {
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
	ProductID   string `json:"product_id" binding:"required" example:"com.fitness.xxx"`       // 產品 id
	OrderID     string `json:"order_id" binding:"required" example:"202105201300687423"`      // 訂單 id
}

type RedeemCourseBody struct {
	OrderID string `json:"order_id" binding:"required" example:"202105201300687423"` // 訂單 id
}

type AppStoreResponseBodyV2 struct {
	SignedPayload string `json:"signedPayload" example:"MIJOlgYJKoZIhvcN..."` // The payload in JSON Web Signature (JWS) format, signed by the App Store
}

type GooglePlayResponseBody struct {
	Message struct {
		Data      string `json:"data"`
		MessageID string `json:"messageId"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type GetSubscriptionsUri struct {
	OriginalTransactionID string `uri:"original_transaction_id" binding:"required" example:"202105201300687423"` // 交易 id
}

type ProductIDUri struct {
	ProductID string `uri:"product_id" binding:"required" example:"5"` // 產品id
}

type GooglePlayAPIGetProductQuery struct {
	PurchaseToken string `form:"purchase_token" binding:"required" example:"d4f5ewq84f65w7e865r1dqw"` // 購買token
}
