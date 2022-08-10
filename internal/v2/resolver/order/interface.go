package order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateCourseOrder(tx *gorm.DB, input *model.APICreateCourseOrderInput) (output model.APICreateCourseOrderOutput)
	APICreateSubscribeOrder(tx *gorm.DB, input *model.APICreateSubscribeOrderInput) (output model.APICreateSubscribeOrderOutput)
	APIGetCMSOrders(input *model.APIGetCMSOrdersInput) (output model.APIGetCMSOrdersOutput)
	APIVerifyAppleReceipt(ctx *gin.Context, tx *gorm.DB, input *model.APIVerifyAppleReceiptInput) (output model.APIVerifyAppleReceiptOutput)
	APIAppStoreNotification(ctx *gin.Context, tx *gorm.DB, input *model.APIAppStoreNotificationInput) (output model.APIAppStoreNotificationOutput)
}
