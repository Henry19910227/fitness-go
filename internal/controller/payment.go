package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
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
	userMidd midd.User,
	courseMidd midd.Course) {

	review := &Payment{PaymentService: PaymentService, CourseService: CourseService}
	baseGroup.POST("/course_order",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		review.CreateCourseOrder)

	baseGroup.POST("/subscribe_order",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		review.CreateSubscribeOrder)

	baseGroup.POST("/verify_apple_receipt",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		review.VerifyAppleReceipt)

	baseGroup.POST("/redeem_course",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		review.RedeemCourse)

	baseGroup.POST("/app_store_notification/v2",
		review.AppStoreNotification)

	baseGroup.POST("/payment_test",
		review.PaymentTest)
}

// CreateCourseOrder 創建課表訂單
// @Summary 創建課表訂單
// @Description 創建課表訂單
// @Tags Payment
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateCourseOrderBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.CourseOrder} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /course_order [POST]
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
// @Tags Payment
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateSubscribeOrderBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.SubscribeOrder} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /subscribe_order [POST]
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
// @Tags Payment
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.VerifyReceiptBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /verify_apple_receipt [POST]
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

// RedeemCourse 兌換免費課表
// @Summary 兌換免費課表
// @Description 兌換免費課表
// @Tags Payment
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.RedeemCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /redeem_course [POST]
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

func (p *Payment) PaymentTest(c *gin.Context) {
	result, err := p.PaymentService.Test(c)
	if err != nil {
		p.JSONErrorResponse(c, err)
		return
	}
	p.JSONSuccessResponse(c, result, "success")
}
