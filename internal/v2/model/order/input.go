package order

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
	PreloadInput
}

type ListInput struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	order_course.CourseIDOptional
	receipt.OriginalTransactionIDOptional
	OrderByInput
	PagingInput
	PreloadInput
}

// APICreateCourseOrderInput /v2/course_order [POST]
type APICreateCourseOrderInput struct {
	UserIDRequired
	Body APICreateCourseOrderBody
}
type APICreateCourseOrderBody struct {
	order_course.CourseIDRequired
}

// APICreateSubscribeOrderInput /v2/subscribe_order [POST]
type APICreateSubscribeOrderInput struct {
	UserIDRequired
	Body APICreateSubscribeOrderBody
}
type APICreateSubscribeOrderBody struct {
	SubscribePlanID int64 `json:"subscribe_plan_id" binding:"required" example:"1"` // 訂閱項目id
}

// APIGetCMSOrdersInput /v2/cms/orders [GET]
type APIGetCMSOrdersInput struct {
	Form APIGetCMSOrdersForm
}
type APIGetCMSOrdersForm struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	PagingInput
	OrderByInput
}

// APIVerifyAppleSubscribeInput /v2/verify_apple_payment [POST]
type APIVerifyAppleSubscribeInput struct {
	UserIDRequired
	Body APIVerifyAppleSubscribeBody
}
type APIVerifyAppleSubscribeBody struct {
	OriginalTransactionID string `json:"original_transaction_id" binding:"required" example:"1000000968276600"` // 初始交易id
}

// APIVerifyAppleReceiptInput /v2/verify_apple_receipt [POST]
type APIVerifyAppleReceiptInput struct {
	UserIDRequired
	Body APIVerifyAppleReceiptBody
}
type APIVerifyAppleReceiptBody struct {
	OrderID     string `json:"order_id" binding:"required" example:"202105201300687423"`      //訂單id
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}

// APIAppStoreNotificationInput /v2/app_store_notification/v2 [POST]
type APIAppStoreNotificationInput struct {
	Body APIAppStoreNotificationBody
}
type APIAppStoreNotificationBody struct {
	SignedPayload string `json:"signedPayload" example:"MIJOlgYJKoZIhvcN..."` // The payload in JSON Web Signature (JWS) format, signed by the App Store
}
