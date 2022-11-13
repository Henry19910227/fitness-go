package order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/required"
	orderCourseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/order_course/optional"
	orderCourseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order_course/required"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	optional.IDField
	PreloadInput
}

type ListInput struct {
	optional.IDField
	optional.UserIDField
	optional.OrderTypeField
	optional.OrderStatusField
	orderCourseOptional.CourseIDField
	receipt.OriginalTransactionIDOptional
	OrderByInput
	PagingInput
	PreloadInput
}

// APICreateCourseOrderInput /v2/course_order [POST]
type APICreateCourseOrderInput struct {
	required.UserIDField
	Body APICreateCourseOrderBody
}
type APICreateCourseOrderBody struct {
	orderCourseRequired.CourseIDField
}

// APICreateSubscribeOrderInput /v2/subscribe_order [POST]
type APICreateSubscribeOrderInput struct {
	required.UserIDField
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
	optional.IDField
	optional.UserIDField
	optional.OrderTypeField
	optional.OrderStatusField
	PagingInput
	OrderByInput
}

// APIVerifyAppleSubscribeInput /v2/verify_apple_payment [POST]
type APIVerifyAppleSubscribeInput struct {
	required.UserIDField
	Body APIVerifyAppleSubscribeBody
}
type APIVerifyAppleSubscribeBody struct {
	OriginalTransactionID string `json:"original_transaction_id" binding:"required" example:"1000000968276600"` // 初始交易id
}

// APIVerifyAppleReceiptInput /v2/verify_apple_receipt [POST]
type APIVerifyAppleReceiptInput struct {
	required.UserIDField
	Body APIVerifyAppleReceiptBody
}
type APIVerifyAppleReceiptBody struct {
	OrderID     string `json:"order_id" binding:"required" example:"202105201300687423"`      //訂單id
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}

// APIVerifyGoogleReceiptInput /v2/verify_google_receipt [POST]
type APIVerifyGoogleReceiptInput struct {
	required.UserIDField
	Body APIVerifyGoogleReceiptBody
}
type APIVerifyGoogleReceiptBody struct {
	OrderID     string `json:"order_id" binding:"required" example:"202105201300687423"`      // 訂單id
	ProductID   string `json:"product_id" binding:"required" example:"com.fitness.xxx"`       // 產品id
	ReceiptData string `json:"receipt_data" binding:"required" example:"MIJOlgYJKoZIhvcN..."` // 收據token
}

// APIAppStoreNotificationInput /v2/app_store_notification/v2 [POST]
type APIAppStoreNotificationInput struct {
	Body APIAppStoreNotificationBody
}
type APIAppStoreNotificationBody struct {
	SignedPayload string `json:"signedPayload" example:"MIJOlgYJKoZIhvcN..."` // The payload in JSON Web Signature (JWS) format, signed by the App Store
}

// APIGooglePlayNotificationInput /v2/google_play_notification [POST]
type APIGooglePlayNotificationInput struct {
	Body APIGooglePlayNotificationBody
}
type APIGooglePlayNotificationBody struct {
	Message struct {
		Data      string `json:"data"`
		MessageID string `json:"messageId"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}
