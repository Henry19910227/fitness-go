package order

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateCourseOrder(ctx *gin.Context)
	CreateSubscribeOrder(ctx *gin.Context)
	UploadAppleSubscribeReceipt(ctx *gin.Context)
	UploadAppleSubscribeReceipts(ctx *gin.Context)
	UploadAppleChargeReceipt(ctx *gin.Context)
	UploadGoogleSubscribeReceipt(ctx *gin.Context)
	UploadGoogleChargeReceipt(ctx *gin.Context)
	VerifyAppleReceipt(ctx *gin.Context)
	VerifyGoogleReceipt(ctx *gin.Context)
	OrderRedeem(ctx *gin.Context)
	AppStoreNotification(ctx *gin.Context)
	GooglePlayNotification(ctx *gin.Context)
	VerifyAppleSubscribe(ctx *gin.Context)
	GetCMSOrders(ctx *gin.Context)

	SyncAppleSubscribeStatusSchedule()
}
