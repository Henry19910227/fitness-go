package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type Payment struct {
	Base
	PaymentService service.Payment
	CourseService  service.Course
}

func NewPayment(baseGroup *gin.RouterGroup,
	PaymentService service.Payment,
	CourseService service.Course,
	userMidd midd.User) {

	payment := &Payment{PaymentService: PaymentService, CourseService: CourseService}
	baseGroup.POST("/course_order",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		payment.CreateCourseOrder)

	baseGroup.POST("/subscribe_order",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		payment.CreateSubscribeOrder)

	baseGroup.POST("/verify_apple_receipt",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		payment.VerifyAppleReceipt)

	baseGroup.POST("/verify_google_receipt",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		payment.VerifyGoogleReceipt)

	baseGroup.POST("/redeem_course",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		payment.RedeemCourse)

	baseGroup.POST("/app_store_notification/v2",
		payment.AppStoreNotification)

	baseGroup.POST("/google_play_notification",
		payment.GooglePlayNotification)

	baseGroup.GET("/app_store/subscriptions/:original_transaction_id",
		payment.GetSubscriptions)

	baseGroup.GET("/app_store/history/:original_transaction_id",
		payment.GetHistory)

	baseGroup.GET("/google_play/access_token",
		payment.GetGooglePlayDeveloperAPIAccessToken)

	baseGroup.GET("/google_play/product/:product_id",
		payment.GetGooglePlayDeveloperAPIProduct)

	baseGroup.GET("/google_play/subscription/:product_id",
		payment.GetGooglePlayDeveloperAPISubscription)

	baseGroup.GET("/app_store/access_token",
		payment.GetAppStoreServerAPIAccessToken)
}

// CreateCourseOrder 創建課表訂單
// @Summary 創建課表訂單
// @Description 創建課表訂單
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateCourseOrderBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.CourseOrder} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/course_order [POST]
func (p *Payment) CreateCourseOrder(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.CreateCourseOrderBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	order, e := p.PaymentService.CreateCourseOrder(c, uid, body.CourseID)
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, order, "success")
}

// CreateSubscribeOrder 創建訂閱訂單
// @Summary 創建訂閱訂單
// @Description 創建訂閱訂單
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateSubscribeOrderBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.SubscribeOrder} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/subscribe_order [POST]
func (p *Payment) CreateSubscribeOrder(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.CreateSubscribeOrderBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	order, e := p.PaymentService.CreateSubscribeOrder(c, uid, body.SubscribePlanID)
	if e != nil {
		p.JSONErrorResponse(c, e)
		return
	}
	p.JSONSuccessResponse(c, order, "success")
}

// VerifyAppleReceipt 驗證apple收據
// @Summary 驗證apple收據
// @Description 驗證apple收據
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.VerifyReceiptBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/verify_apple_receipt [POST]
func (p *Payment) VerifyAppleReceipt(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.VerifyReceiptBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.PaymentService.VerifyAppleReceipt(c, uid, body.OrderID, body.ReceiptData); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, nil, "success")
}

// VerifyGoogleReceipt 驗證google收據
// @Summary 驗證google收據
// @Description 驗證google收據
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.VerifyGoogleReceiptBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/verify_google_receipt [POST]
func (p *Payment) VerifyGoogleReceipt(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.VerifyGoogleReceiptBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.PaymentService.VerifyGoogleReceipt(c, uid, body.OrderID, body.ProductID, body.ReceiptData); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, nil, "success")
}

// RedeemCourse 兌換免費課表
// @Summary 兌換免費課表
// @Description 兌換免費課表
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.RedeemCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/redeem_course [POST]
func (p *Payment) RedeemCourse(c *gin.Context) {
	uid, err := p.GetUID(c)
	if err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.RedeemCourseBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := p.PaymentService.VerifyFreeCourseOrder(c, uid, body.OrderID); err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, nil, "success")
}

func (p *Payment) AppStoreNotification(c *gin.Context) {
	var body validator.AppStoreResponseBodyV2
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	err := p.PaymentService.HandleAppStoreNotification(c, body.SignedPayload)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, nil, "success")
}

func (p *Payment) GooglePlayNotification(c *gin.Context) {
	var body validator.GooglePlayResponseBody
	if err := c.ShouldBindJSON(&body); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	err := p.PaymentService.HandleGooglePlayNotification(c, body.Message.Data)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, nil, "success")
}

// GetSubscriptions 獲取訂閱資料(測試用)
// @Summary 獲取訂閱資料(測試用)
// @Description 獲取訂閱資料(測試用)
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Param original_transaction_id path string true "交易id"
// @Success 200 {object} model.SuccessResult{data=dto.IAPSubscribeAPIResponse} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/app_store/subscriptions/{original_transaction_id} [GET]
func (p *Payment) GetSubscriptions(c *gin.Context) {
	var uri validator.GetSubscriptionsUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := p.PaymentService.GetAppStoreAPISubscriptions(c, uri.OriginalTransactionID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, result, "success")
}

// GetHistory 獲取交易歷史資料(測試用)
// @Summary 獲取交易歷史資料(測試用)
// @Description 獲取交易歷史資料(測試用)
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Param original_transaction_id path string true "交易id"
// @Success 200 {object} model.SuccessResult{data=dto.IAPSubscribeAPIResponse} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/app_store/history/{original_transaction_id} [GET]
func (p *Payment) GetHistory(c *gin.Context) {
	var uri validator.GetSubscriptionsUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := p.PaymentService.GetAppStoreAPIHistory(c, uri.OriginalTransactionID)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, result, "success")
}

// GetAppStoreServerAPIAccessToken 取得 App Store Server api access token
// @Summary 取得 Apple Store api access token
// @Description 取得 Apple Store api access token
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Success 200 {object} model.SuccessResult "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/app_store/access_token [GET]
func (p *Payment) GetAppStoreServerAPIAccessToken(c *gin.Context) {
	accessToken, err := p.PaymentService.GetAppleStoreApiAccessToken(c)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, accessToken, "success")
}

// GetGooglePlayDeveloperAPIAccessToken 取得 google play developer api access token(測試用)
// @Summary 取得 google play api access token(測試用)
// @Description 取得 google play api access token(測試用)
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Success 200 {object} model.SuccessResult "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/google_play/access_token [GET]
func (p *Payment) GetGooglePlayDeveloperAPIAccessToken(c *gin.Context) {
	accessToken, err := p.PaymentService.GetGooglePlayApiAccessToken(c)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, accessToken, "success")
}

// GetGooglePlayDeveloperAPIProduct 取得 google play developer api product(測試用)
// @Summary 取得 google play developer api product(測試用)
// @Description 取得 google play developer api product(測試用)
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Param product_id path string true "產品id"
// @Param purchase_token query string true "收據token"
// @Success 200 {object} model.SuccessResult{data=dto.IABProductAPIResponse} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/google_play/product/{product_id} [GET]
func (p *Payment) GetGooglePlayDeveloperAPIProduct(c *gin.Context) {
	var uri validator.ProductIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var query validator.GooglePlayAPIGetProductQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := p.PaymentService.GetGooglePlayAPIProduct(c, uri.ProductID, query.PurchaseToken)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, result, "success")
}

// GetGooglePlayDeveloperAPISubscription 取得 google play developer api subscription(測試用)
// @Summary 取得 google play developer api subscription(測試用)
// @Description 取得 google play developer api subscription(測試用)
// @Tags Payment_v1
// @Accept json
// @Produce json
// @Param product_id path string true "產品id"
// @Param purchase_token query string true "收據token"
// @Success 200 {object} model.SuccessResult{data=dto.IABSubscriptionAPIResponse} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/google_play/subscription/{product_id} [GET]
func (p *Payment) GetGooglePlayDeveloperAPISubscription(c *gin.Context) {
	var uri validator.ProductIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var query validator.GooglePlayAPIGetProductQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		p.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := p.PaymentService.GetGooglePlayAPISubscription(c, uri.ProductID, query.PurchaseToken)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, result, "success")
}
