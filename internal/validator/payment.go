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

type RedeemCourseBody struct {
	OrderID string `json:"order_id" binding:"required" example:"202105201300687423"` // 訂單 id
}

type AppStoreResponseBodyV2 struct {
	SignedPayload string `json:"signedPayload" example:"MIJOlgYJKoZIhvcN..."` // The payload in JSON Web Signature (JWS) format, signed by the App Store
}

type GetSubscriptionsUri struct {
	OriginalTransactionID string `uri:"original_transaction_id" binding:"required" example:"202105201300687423"` // 交易 id
}
