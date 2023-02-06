package order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_create_subscribe_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_get_cms_user_orders"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_order_redeem"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_charge_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_subscribe_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_subscribe_receipts"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_google_charge_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_google_subscribe_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_google_subscribe_receipts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSUserOrders(input *api_get_cms_user_orders.Input) (output api_get_cms_user_orders.Output)

	APICreateCourseOrder(tx *gorm.DB, input *model.APICreateCourseOrderInput) (output model.APICreateCourseOrderOutput)
	APICreateSubscribeOrder(tx *gorm.DB, input *api_create_subscribe_order.Input) (output api_create_subscribe_order.Output)
	APIGetCMSOrders(input *model.APIGetCMSOrdersInput) (output model.APIGetCMSOrdersOutput)
	APIVerifyAppleReceipt(ctx *gin.Context, tx *gorm.DB, input *model.APIVerifyAppleReceiptInput) (output model.APIVerifyAppleReceiptOutput)
	APIUploadAppleSubscribeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_apple_subscribe_receipt.Input) (output api_upload_apple_subscribe_receipt.Output)
	APIUploadAppleSubscribeReceipts(ctx *gin.Context, input *api_upload_apple_subscribe_receipts.Input) (output api_upload_apple_subscribe_receipts.Output)
	APIUploadAppleChargeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_apple_charge_receipt.Input) (output api_upload_apple_charge_receipt.Output)
	APIUploadGoogleSubscribeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_google_subscribe_receipt.Input) (output api_upload_google_subscribe_receipt.Output)
	APIUploadGoogleSubscribeReceipts(ctx *gin.Context, input *api_upload_google_subscribe_receipts.Input) (output api_upload_google_subscribe_receipts.Output)
	APIUploadGoogleChargeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_google_charge_receipt.Input) (output api_upload_google_charge_receipt.Output)
	APIVerifyGoogleReceipt(ctx *gin.Context, tx *gorm.DB, input *model.APIVerifyGoogleReceiptInput) (output model.APIVerifyGoogleReceiptOutput)
	APIOrderRedeem(tx *gorm.DB, input *api_order_redeem.Input) (output api_order_redeem.Output)
	APIAppStoreNotification(ctx *gin.Context, tx *gorm.DB, input *model.APIAppStoreNotificationInput) (output model.APIAppStoreNotificationOutput)
	APIGooglePlayNotification(ctx *gin.Context, tx *gorm.DB, input *model.APIGooglePlayNotificationInput) (output model.APIGooglePlayNotificationOutput)
	APIVerifyAppleSubscribe(input *model.APIVerifyAppleSubscribeInput) (output model.APIVerifyAppleSubscribeOutput)

	SyncAppleSubscribeStatusSchedule(tx *gorm.DB)
}
