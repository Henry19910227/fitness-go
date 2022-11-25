package order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_create_subscribe_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_order_redeem"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_charge_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_subscribe_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_subscribe_receipts"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_google_charge_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_google_subscribe_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver order.Resolver
}

func New(resolver order.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateCourseOrder 創建課表訂單
// @Summary 創建課表訂單
// @Description 創建課表訂單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APICreateCourseOrderBody true "輸入參數"
// @Success 200 {object} order.APICreateCourseOrderOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/course_order [POST]
func (c *controller) CreateCourseOrder(ctx *gin.Context) {
	input := model.APICreateCourseOrderInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateCourseOrder(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// CreateSubscribeOrder 創建訂閱訂單
// @Summary 創建訂閱訂單
// @Description 創建訂閱訂單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_create_subscribe_order.Body true "輸入參數"
// @Success 200 {object} api_create_subscribe_order.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/subscribe_order [POST]
func (c *controller) CreateSubscribeOrder(ctx *gin.Context) {
	input := api_create_subscribe_order.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateSubscribeOrder(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UploadAppleSubscribeReceipt 上傳單張apple訂閱收據
// @Summary 上傳單張apple訂閱收據
// @Description 上傳單張apple訂閱收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_upload_apple_subscribe_receipt.Body true "輸入參數"
// @Success 200 {object} api_upload_apple_subscribe_receipt.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/apple_subscribe_receipt [POST]
func (c *controller) UploadAppleSubscribeReceipt(ctx *gin.Context) {
	input := api_upload_apple_subscribe_receipt.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUploadAppleSubscribeReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UploadAppleSubscribeReceipts 上傳多張apple訂閱收據
// @Summary 上傳多張apple訂閱收據
// @Description 上傳多張apple訂閱收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_upload_apple_subscribe_receipts.Body true "輸入參數"
// @Success 200 {object} api_upload_apple_subscribe_receipts.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/apple_subscribe_receipts [POST]
func (c *controller) UploadAppleSubscribeReceipts(ctx *gin.Context) {
	input := api_upload_apple_subscribe_receipts.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUploadAppleSubscribeReceipts(ctx, &input)
	ctx.JSON(http.StatusOK, output)
}

// UploadAppleChargeReceipt 上傳apple付費收據
// @Summary 上傳apple付費收據
// @Description 上傳apple付費收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_upload_apple_charge_receipt.Body true "輸入參數"
// @Success 200 {object} api_upload_apple_charge_receipt.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/apple_charge_receipt [POST]
func (c *controller) UploadAppleChargeReceipt(ctx *gin.Context) {
	input := api_upload_apple_charge_receipt.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUploadAppleChargeReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UploadGoogleSubscribeReceipt 上傳單張google訂閱收據
// @Summary 上傳單張google訂閱收據
// @Description 上傳單張google訂閱收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_upload_google_subscribe_receipt.Body true "輸入參數"
// @Success 200 {object} api_upload_google_subscribe_receipt.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/google_subscribe_receipt [POST]
func (c *controller) UploadGoogleSubscribeReceipt(ctx *gin.Context) {
	input := api_upload_google_subscribe_receipt.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUploadGoogleSubscribeReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UploadGoogleChargeReceipt 上傳google付費收據
// @Summary 上傳google付費收據
// @Description 上傳google付費收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_upload_google_charge_receipt.Body true "輸入參數"
// @Success 200 {object} api_upload_google_charge_receipt.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/google_charge_receipt [POST]
func (c *controller) UploadGoogleChargeReceipt(ctx *gin.Context) {
	input := api_upload_google_charge_receipt.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUploadGoogleChargeReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// VerifyAppleReceipt 驗證apple收據
// @Summary 驗證apple收據
// @Description 驗證apple收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APIVerifyAppleReceiptBody true "輸入參數"
// @Success 200 {object} order.APIVerifyAppleReceiptOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/verify_apple_receipt [POST]
func (c *controller) VerifyAppleReceipt(ctx *gin.Context) {
	input := model.APIVerifyAppleReceiptInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIVerifyAppleReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// VerifyGoogleReceipt 驗證google收據
// @Summary 驗證google收據
// @Description 驗證google收據
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APIVerifyGoogleReceiptBody true "輸入參數"
// @Success 200 {object} order.APIVerifyGoogleReceiptOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/verify_google_receipt [POST]
func (c *controller) VerifyGoogleReceipt(ctx *gin.Context) {
	input := model.APIVerifyGoogleReceiptInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIVerifyGoogleReceipt(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// OrderRedeem 訂單兌換免費課表
// @Summary 訂單兌換免費課表
// @Description 訂單兌換免費課表
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_id path string true "訂單ID"
// @Success 200 {object} api_order_redeem.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/order/{order_id}/redeem [POST]
func (c *controller) OrderRedeem(ctx *gin.Context) {
	input := api_order_redeem.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIOrderRedeem(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// AppStoreNotification app store 訂閱週期通知
// @Summary app store 訂閱週期通知
// @Description app store 訂閱週期通知
// @Tags 支付通知_v2
// @Accept json
// @Produce json
// @Param json_body body order.APIAppStoreNotificationBody true "輸入參數"
// @Success 200 {object} order.APIAppStoreNotificationOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/app_store_notification/v2 [POST]
func (c *controller) AppStoreNotification(ctx *gin.Context) {
	input := model.APIAppStoreNotificationInput{}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIAppStoreNotification(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	if output.Code != 0 {
		ctx.JSON(http.StatusBadRequest, output)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

// GooglePlayNotification google play 訂閱週期通知
// @Summary app store 訂閱週期通知
// @Description app store 訂閱週期通知
// @Tags 支付通知_v2
// @Accept json
// @Produce json
// @Param json_body body order.APIGooglePlayNotificationBody true "輸入參數"
// @Success 200 {object} order.APIGooglePlayNotificationOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/google_play_notification [POST]
func (c *controller) GooglePlayNotification(ctx *gin.Context) {
	input := model.APIGooglePlayNotificationInput{}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGooglePlayNotification(ctx, ctx.MustGet("tx").(*gorm.DB), &input)
	if output.Code != 0 {
		ctx.JSON(http.StatusBadRequest, output)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

// VerifyAppleSubscribe 驗證帳戶是否允許訂閱
// @Summary 驗證帳戶是否允許訂閱
// @Description 在創建訂閱訂單前，確認該帳戶是否可訂閱的API
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APIVerifyAppleSubscribeBody true "輸入參數"
// @Success 200 {object} order.APIVerifyAppleSubscribeOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/verify_apple_subscribe [POST]
func (c *controller) VerifyAppleSubscribe(ctx *gin.Context) {
	input := model.APIVerifyAppleSubscribeInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIVerifyAppleSubscribe(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSOrders 獲取訂單列表
// @Summary 獲取訂單列表
// @Description 獲取訂單列表
// @Tags CMS訂單管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_id query string false "訂單ID"
// @Param user_id query int64 false "用戶ID"
// @Param type query int false "訂單類型(1:課表購買/2:會員訂閱)"
// @Param order_status query int false "訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} order.APIGetCMSOrdersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/orders [GET]
func (c *controller) GetCMSOrders(ctx *gin.Context) {
	input := model.APIGetCMSOrdersInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSOrders(&input)
	ctx.JSON(http.StatusOK, output)
}

func (c *controller) SyncAppleSubscribeStatusSchedule() {
	txHandle := orm.Shared().DB().Begin()
	defer txHandle.Rollback()
	c.resolver.SyncAppleSubscribeStatusSchedule(txHandle)
}
