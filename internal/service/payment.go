package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"time"
)

type payment struct {
	Base
	userRepo            repository.User
	orderRepo           repository.Order
	saleRepo            repository.Sale
	subscribePlanRepo   repository.SubscribePlan
	courseRepo          repository.Course
	receiptRepo         repository.Receipt
	userCourseAssetRepo repository.UserCourseAsset
	subscribeLogRepo    repository.SubscribeLog
	purchaseLogRepo     repository.PurchaseLog
	subscribeInfo       repository.UserSubscribeInfo
	transactionRepo     repository.Transaction
	iapHandler          handler.IAP
	reqTool             tool.HttpRequest
	jwtTool             tool.JWT
	errHandler          errcode.Handler
}

func NewPayment(userRepo repository.User, orderRepo repository.Order, saleRepo repository.Sale, subscribePlanRepo repository.SubscribePlan,
	courseRepo repository.Course, receiptRepo repository.Receipt,
	purchaseRepo repository.UserCourseAsset, subscribeLogRepo repository.SubscribeLog,
	purchaseLogRepo repository.PurchaseLog, memberRepo repository.UserSubscribeInfo,
	transactionRepo repository.Transaction, iapHandler handler.IAP, reqTool tool.HttpRequest,
	jwtTool tool.JWT, errHandler errcode.Handler) Payment {
	return &payment{userRepo: userRepo, orderRepo: orderRepo, saleRepo: saleRepo, subscribePlanRepo: subscribePlanRepo,
		courseRepo: courseRepo, receiptRepo: receiptRepo,
		userCourseAssetRepo: purchaseRepo, subscribeLogRepo: subscribeLogRepo, purchaseLogRepo: purchaseLogRepo,
		subscribeInfo: memberRepo, transactionRepo: transactionRepo,
		reqTool: reqTool, jwtTool: jwtTool, iapHandler: iapHandler, errHandler: errHandler}
}

func (p *payment) GetSubscriptions(c *gin.Context, originalTransactionID string) (*dto.IAPSubscribeResponse, errcode.Error) {
	result, err := p.iapHandler.GetSubscriptionAPI(originalTransactionID)
	if err != nil {
		return nil, p.errHandler.Set(c, "iap handler", err)
	}
	return result, nil
}

func (p *payment) CreateCourseOrder(c *gin.Context, uid int64, courseID int64) (*dto.CourseOrder, errcode.Error) {
	//檢查是此課表是否已購買
	courseAsset, err := p.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
		UserID:   uid,
		CourseID: courseID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	if courseAsset != nil {
		if courseAsset.Available == 1 {
			return nil, p.errHandler.Custom(8999, errors.New("此課表已被購買，無法再創建訂單"))
		}
	}
	//檢查是否有尚未付款的相同訂單
	orderData, err := p.orderRepo.FindOrderByCourseID(uid, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	if orderData != nil {
		return parserCourseOrder(orderData), nil
	}
	//建立新的訂單
	course, err := p.courseRepo.FindCourseProduct(courseID)
	if err != nil {
		return nil, p.errHandler.Set(c, "course repo", err)
	}
	if course.CourseStatus != int(global.Sale) {
		return nil, p.errHandler.Custom(8999, errors.New("必須為銷售狀態的課表才可被加入訂單"))
	}
	if course.SaleType != int(global.SaleTypeFree) && course.SaleType != int(global.SaleTypeCharge) {
		return nil, p.errHandler.Custom(8999, errors.New("商品必須為免費課表或付費課表類型才可創建此訂單"))
	}
	if course.SaleType != int(global.SaleTypeFree) && course.SaleID == nil {
		return nil, p.errHandler.Custom(8999, errors.New("付費課表必須有 sale id"))
	}
	// 創建訂單
	orderID, err := p.orderRepo.CreateCourseOrder(&model.CreateOrderParam{
		UserID:     uid,
		SaleItemID: course.SaleID,
		CourseID:   courseID,
	})
	if err != nil {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	// parser 資料
	data, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	return parserCourseOrder(data), nil
}

func (p *payment) CreateSubscribeOrder(c *gin.Context, uid int64, subscribePlanID int64) (*dto.SubscribeOrder, errcode.Error) {
	// 查詢用戶訂閱資訊
	subscribeInfo, err := p.subscribeInfo.FindSubscribeInfo(uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, p.errHandler.Set(c, "user subscribe repo", err)
	}
	if subscribeInfo == nil {
		// 查找訂閱方案
		subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByID(subscribePlanID)
		if err != nil {
			return nil, p.errHandler.Set(c, "order repo", err)
		}
		//創建新的訂單
		orderID, err := p.orderRepo.CreateSubscribeOrder(&model.CreateSubscribeOrderParam{
			UserID:          uid,
			SubscribePlanID: subscribePlan.ID,
		})
		if err != nil {
			return nil, p.errHandler.Set(c, "order repo", err)
		}
		// 查找剛創建的訂單
		data, err := p.orderRepo.FindOrder(orderID)
		if err != nil {
			return nil, p.errHandler.Set(c, "order repo", err)
		}
		return parserSubscribeOrder(data), nil
	}
	if subscribeInfo.OrderID == nil {
		return nil, p.errHandler.Custom(8999, errors.New("格式錯誤"))
	}
	//查詢訂單收據
	receipts, err := p.receiptRepo.FindReceiptsByOrderID(*subscribeInfo.OrderID, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return nil, p.errHandler.Set(c, "receipt repo", err)
	}
	if len(receipts) == 0 {
		// 有訂單但無收據的情況下返回最新創建的訂單
		data, err := p.orderRepo.FindOrder(*subscribeInfo.OrderID)
		if err != nil {
			return nil, p.errHandler.Set(c, "order repo", err)
		}
		return parserSubscribeOrder(data), nil
	}
	if global.PaymentType(receipts[0].PaymentType) == global.ApplePaymentType {
		response, err := p.iapHandler.GetSubscriptionAPI(receipts[0].OriginalTransactionID)
		if err != nil {
			return nil, p.errHandler.Set(c, "iap handler", err)
		}
		if response.Status == 2 { // 當前訂閱已過期
			// 查找剛創建的訂單
			data, err := p.orderRepo.FindOrder(*subscribeInfo.OrderID)
			if err != nil {
				return nil, p.errHandler.Set(c, "order repo", err)
			}
			return parserSubscribeOrder(data), nil
		}
		return nil, p.errHandler.Custom(8999, errors.New("目前已經是訂閱會員"))
	}
	if global.PaymentType(receipts[0].PaymentType) == global.GooglePaymentType {
		return nil, p.errHandler.Custom(8999, errors.New("尚未支援"))
	}
	return nil, p.errHandler.Custom(8999, errors.New("創建訂單失敗"))
}

func (p *payment) CheckSubscribeStatus(c *gin.Context, uid int64, orderID string) bool {
	return false
}

func (p *payment) VerifyFreeCourseOrder(c *gin.Context, uid int64, orderID string) errcode.Error {
	//取得訂單資訊
	order, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	if order.UserID != uid {
		return p.errHandler.Custom(8999, errors.New("無效的收據(此訂單非本人)"))
	}
	if global.PaymentOrderType(order.OrderType) != global.BuyCourseOrderType {
		return p.errHandler.Custom(8999, errors.New("此訂單類型錯誤"))
	}
	if order.OrderCourse == nil {
		return p.errHandler.Custom(8999, errors.New("此訂單的未註名商品"))
	}
	//處理課表購買訂單
	if order.OrderCourse.Course.SaleType != int(global.SaleTypeFree) {
		return p.errHandler.Custom(8999, errors.New("此訂單的商品非免費"))
	}
	//驗證訂單狀態
	if global.OrderStatus(order.OrderStatus) != global.PendingOrderStatus {
		return p.errHandler.Custom(8999, errors.New("此訂單已失效"))
	}
	if err := p.handleFreeCourseTrade(c, parserCourseOrder(order)); err != nil {
		return err
	}
	return nil
}

func (p *payment) VerifyAppleReceipt(c *gin.Context, uid int64, orderID string, receiptData string) errcode.Error {
	//取得訂單資訊
	order, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	if order.UserID != uid {
		return p.errHandler.Custom(8999, errors.New("無效的收據(此訂單非本人)"))
	}
	//apple server 正式區驗證收據
	param := map[string]interface{}{
		"receipt-data":             receiptData,
		"password":                 p.iapHandler.Password(),
		"exclude-old-transactions": 1,
	}
	result, err := p.reqTool.SendPostRequestWithJsonBody(p.iapHandler.ProductURL(), param)
	if err != nil {
		return p.errHandler.Set(c, "req tool", err)
	}
	var response dto.AppleReceiptResponse
	if err := p.iapHandler.ParserAppleReceipt(result, &response); err != nil {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	//apple server 測試區驗證收據
	if response.Status == 21007 {
		result, err := p.reqTool.SendPostRequestWithJsonBody(p.iapHandler.SandboxURL(), param)
		if err != nil {
			return p.errHandler.Set(c, "req tool", err)
		}
		if err := p.iapHandler.ParserAppleReceipt(result, &response); err != nil {
			return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
		}
	}
	//驗證收據結果
	if response.Status != 0 {
		//更新訂單狀態
		_ = p.orderRepo.UpdateOrderStatus(nil, order.ID, global.ErrorOrderStatus)
		return p.errHandler.Custom(8999, errors.New("收據驗證錯誤"))
	}
	//處理課表購買訂單
	if order.OrderType == int(global.BuyCourseOrderType) {
		if err := p.handleBuyCourseTrade(c, parserCourseOrder(order), receiptData, &response); err != nil {
			return err
		}
	}
	//處理會員訂閱訂單
	if order.OrderType == int(global.SubscribeOrderType) {
		if err := p.handleSubscribeTrade(c, uid, parserSubscribeOrder(order), receiptData, &response); err != nil {
			return err
		}
	}
	return nil
}

func (p *payment) VerifyGoogleReceipt(c *gin.Context, uid int64, orderID string, receiptData string) errcode.Error {
	//取得訂單資訊
	order, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	if order.UserID != uid {
		return p.errHandler.Custom(8999, errors.New("無效的收據(此訂單非本人)"))
	}
	response := dto.GoogleReceiptResponse{
		OrderID: "TEST-" + time.Now().Format("20060102150405") + strconv.Itoa(int(randRange(100, 999))),
	}
	//處理課表購買訂單
	if order.OrderType == int(global.BuyCourseOrderType) {
		if err := p.handleBuyCourseTradeForGoogle(c, parserCourseOrder(order), receiptData, &response); err != nil {
			return err
		}
	}
	//處理會員訂閱訂單
	if order.OrderType == int(global.SubscribeOrderType) {
		if err := p.handleSubscribeTradeForGoogle(c, uid, parserSubscribeOrder(order), receiptData, &response); err != nil {
			return err
		}
	}
	panic("implement me")
}

func (p *payment) HandleAppStoreNotification(c *gin.Context, base64PayloadString string) errcode.Error {
	// 解析字串
	response, err := p.iapHandler.DecodeIAPNotificationResponse(base64PayloadString)
	if err != nil {
		return p.errHandler.Set(c, "iap parser", err)
	}
	subscribeLogType := p.iapHandler.ParserIAPNotificationType(response.NotificationType, response.Subtype)
	// 存取log
	_, err = p.subscribeLogRepo.SaveSubscribeLog(nil, &model.CreateSubscribeLogParam{
		OriginalTransactionID: response.Data.SignedTransactionInfo.OriginalTransactionId,
		TransactionID:         response.Data.SignedTransactionInfo.TransactionId,
		PurchaseDate:          parserIAPDate(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"),
		ExpiresDate:           parserIAPDate(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"),
		Type:                  string(subscribeLogType),
		Msg:                   fmt.Sprintf("%s %s", response.NotificationType, response.Subtype),
	})
	if err != nil {
		return p.errHandler.Set(c, "subscribe log repo", err)
	}
	// 查詢用戶訂閱資料
	info, err := p.subscribeInfo.FindSubscribeInfoByOriginalTransactionID(response.Data.SignedTransactionInfo.OriginalTransactionId)
	if err != nil {
		return p.errHandler.Set(c, "subscribe info repo", err)
	}
	if info.OrderID == nil {
		return p.errHandler.Custom(8999, errors.New("缺少 order id 資訊"))
	}
	// 驗證訂單初始交易id
	receipts, err := p.receiptRepo.FindReceiptsByOrderID(*info.OrderID, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return p.errHandler.Set(c, "receipt repo", err)
	}
	if len(receipts) == 0 {
		return p.errHandler.Custom(8999, errors.New("訂單缺少收據資訊"))
	}
	if receipts[0].OriginalTransactionID != response.Data.SignedTransactionInfo.OriginalTransactionId {
		return p.errHandler.Custom(8999, errors.New("訂單初始交易id不一致"))
	}
	tx := p.transactionRepo.CreateTransaction()
	// 存取收據
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               *info.OrderID,
		PaymentType:           int(global.ApplePaymentType),
		ReceiptToken:          "",
		OriginalTransactionID: response.Data.SignedTransactionInfo.OriginalTransactionId,
		TransactionID:         response.Data.SignedTransactionInfo.TransactionId,
		ProductID:             response.Data.SignedTransactionInfo.ProductId,
		Quantity:              1,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	// 查詢當前訂閱項目
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(response.Data.SignedTransactionInfo.ProductId)
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	// 修改訂單訂閱項目(升級or降級狀態)
	if err := p.orderRepo.UpdateOrderSubscribePlan(tx, *info.OrderID, subscribePlan.ID); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	// 修改用戶訂閱資訊
	var subscribeStatus = global.ValidSubscribeStatus
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		subscribeStatus = global.NoneSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          info.UserID,
		OrderID:         *info.OrderID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       parserIAPDate(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"),
		ExpiresDate:     parserIAPDate(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "subscribe info repo", err)
	}
	//更新會員類型
	var userType = global.SubscribeUserType
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		userType = global.NormalUserType
	}
	ut := int(userType)
	if err := p.userRepo.UpdateUserByUID(tx, info.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	p.transactionRepo.FinishTransaction(tx)
	fmt.Printf("NotificationType: %v \n", response.NotificationType)
	fmt.Printf("Subtype: %v \n", response.Subtype)

	fmt.Printf("Data.Environment: %v \n", response.Data.Environment)

	fmt.Printf("SignedRenewalInfo.AutoRenewProductId: %v \n", response.Data.SignedRenewalInfo.AutoRenewProductId)
	fmt.Printf("SignedRenewalInfo.AutoRenewStatus: %v \n", response.Data.SignedRenewalInfo.AutoRenewStatus)
	fmt.Printf("SignedRenewalInfo.ExpirationIntent: %v \n", response.Data.SignedRenewalInfo.ExpirationIntent)
	fmt.Printf("SignedRenewalInfo.GracePeriodExpiresDate: %v \n", parserIAPDate(response.Data.SignedRenewalInfo.GracePeriodExpiresDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedRenewalInfo.IsInBillingRetryPeriod: %v \n", response.Data.SignedRenewalInfo.IsInBillingRetryPeriod)
	fmt.Printf("SignedRenewalInfo.OfferIdentifier: %v \n", response.Data.SignedRenewalInfo.OfferIdentifier)
	fmt.Printf("SignedRenewalInfo.OfferType: %v \n", response.Data.SignedRenewalInfo.OfferType)
	fmt.Printf("SignedRenewalInfo.OriginalTransactionId: %v \n", response.Data.SignedRenewalInfo.OriginalTransactionId)
	fmt.Printf("SignedRenewalInfo.PriceIncreaseStatus: %v \n", response.Data.SignedRenewalInfo.PriceIncreaseStatus)
	fmt.Printf("SignedRenewalInfo.ProductId: %v \n", response.Data.SignedRenewalInfo.ProductId)
	fmt.Printf("SignedRenewalInfo.SignedDate: %v \n", parserIAPDate(response.Data.SignedRenewalInfo.SignedDate/1000).Format("2006-01-02 15:04:05"))

	fmt.Printf("SignedTransactionInfo.AppAccountToken: %v \n", response.Data.SignedTransactionInfo.AppAccountToken)
	fmt.Printf("SignedTransactionInfo.BundleId: %v \n", response.Data.SignedTransactionInfo.BundleId)
	fmt.Printf("SignedTransactionInfo.ExpiresDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.ExpiresDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.InAppOwnershipType: %v \n", response.Data.SignedTransactionInfo.InAppOwnershipType)
	fmt.Printf("SignedTransactionInfo.IsUpgraded: %v \n", response.Data.SignedTransactionInfo.IsUpgraded)
	fmt.Printf("SignedTransactionInfo.OfferIdentifier: %v \n", response.Data.SignedTransactionInfo.OfferIdentifier)
	fmt.Printf("SignedTransactionInfo.OfferType: %v \n", response.Data.SignedTransactionInfo.OfferType)
	fmt.Printf("SignedTransactionInfo.OriginalPurchaseDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.OriginalPurchaseDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.OriginalTransactionId: %v \n", response.Data.SignedTransactionInfo.OriginalTransactionId)
	fmt.Printf("SignedTransactionInfo.ProductId: %v \n", response.Data.SignedTransactionInfo.ProductId)
	fmt.Printf("SignedTransactionInfo.PurchaseDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.PurchaseDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.Quantity: %v \n", response.Data.SignedTransactionInfo.Quantity)
	fmt.Printf("SignedTransactionInfo.RevocationDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.RevocationDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.RevocationReason: %v \n", response.Data.SignedTransactionInfo.RevocationReason)
	fmt.Printf("SignedTransactionInfo.SignedDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.SignedDate/1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.SubscriptionGroupIdentifier: %v \n", response.Data.SignedTransactionInfo.SubscriptionGroupIdentifier)
	fmt.Printf("SignedTransactionInfo.TransactionId: %v \n", response.Data.SignedTransactionInfo.TransactionId)
	fmt.Printf("SignedTransactionInfo.Type: %v \n", response.Data.SignedTransactionInfo.Type)
	fmt.Printf("SignedTransactionInfo.WebOrderLineItemId: %v \n", response.Data.SignedTransactionInfo.WebOrderLineItemId)

	return nil
}

func (p *payment) handleFreeCourseTrade(c *gin.Context, order *dto.CourseOrder) errcode.Error {
	if err := p.handleCourseOrderTrade(order, global.NonePaymentType, "",
		"", "", 1, ""); err != nil {
		return p.errHandler.Set(c, "trade error", err)
	}
	return nil
}

func (p *payment) handleBuyCourseTrade(c *gin.Context, order *dto.CourseOrder, receiptData string, response *dto.AppleReceiptResponse) errcode.Error {
	//驗證產品id
	if order.SaleItem.ProductID != response.Receipt.InApp[0].ProductID {
		return p.errHandler.Custom(8999, errors.New("無效的收據(產品ID不匹配)"))
	}
	item := response.Receipt.InApp[0]
	quantity, err := strconv.Atoi(item.Quantity)
	if err != nil {
		return p.errHandler.Set(c, "Atoi error", err)
	}
	if err := p.handleCourseOrderTrade(order, global.ApplePaymentType, item.OriginalTransactionID,
		item.TransactionID, item.ProductID, quantity, receiptData); err != nil {
		return p.errHandler.Set(c, "trade error", err)
	}
	return nil
}

func (p *payment) handleBuyCourseTradeForGoogle(c *gin.Context, order *dto.CourseOrder, receiptData string, response *dto.GoogleReceiptResponse) errcode.Error {
	if err := p.handleCourseOrderTrade(order, global.GooglePaymentType, response.OrderID,
		response.OrderID, order.SaleItem.ProductID, 1, receiptData); err != nil {
		return p.errHandler.Set(c, "trade error", err)
	}
	return nil
}

func (p *payment) handleCourseOrderTrade(order *dto.CourseOrder, paymentType global.PaymentType,
	originalTransactionID string, transactionID string, productID string, quantity int, receiptData string) error {
	//創建transaction
	tx := p.transactionRepo.CreateTransaction()
	//存入收據
	_, err := p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
		PaymentType:           int(paymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: originalTransactionID,
		TransactionID:         transactionID,
		ProductID:             productID,
		Quantity:              quantity,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	//存入購買Log
	_, err = p.purchaseLogRepo.CreatePurchaseLog(tx, &model.CreatePurchaseLogParam{
		UserID:  order.UserID,
		OrderID: order.ID,
		Type:    global.BuyPurchaseLogType,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	//存入課表購買紀錄
	_, err = p.userCourseAssetRepo.CreateUserCourseAsset(tx, &model.CreateUserCourseAssetParam{
		UserID:   order.UserID,
		CourseID: order.Course.ID,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, order.ID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return err
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func (p *payment) handleSubscribeTrade(c *gin.Context, uid int64, order *dto.SubscribeOrder, receiptData string, response *dto.AppleReceiptResponse) errcode.Error {
	//初始化當前訂單id
	currentOrderID := order.ID
	//驗證收據格式
	if len(response.LatestReceiptInfo) == 0 || len(response.PendingRenewalInfo) == 0 {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	//查詢用戶訂閱資訊
	subscribeInfo, err := p.subscribeInfo.FindSubscribeInfo(uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return p.errHandler.Set(c, "user subscribe repo", err)
	}
	if subscribeInfo != nil {
		if subscribeInfo.OrderID != nil {
			currentOrderID = *subscribeInfo.OrderID
		}
	}
	//驗證訂單初始交易id
	receipts, err := p.receiptRepo.FindReceiptsByOrderID(currentOrderID, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return p.errHandler.Set(c, "receipt repo", err)
	}
	if len(receipts) > 0 {
		//驗證是否是相同的初始id訂單
		if response.LatestReceiptInfo[0].OriginalTransactionID != receipts[0].OriginalTransactionID {
			// 查找訂閱方案
			subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(response.LatestReceiptInfo[0].ProductID)
			if err != nil {
				return p.errHandler.Set(c, "order repo", err)
			}
			//創建新的訂單
			newOrderID, err := p.orderRepo.CreateSubscribeOrder(&model.CreateSubscribeOrderParam{
				UserID:          uid,
				SubscribePlanID: subscribePlan.ID,
			})
			if err != nil {
				return p.errHandler.Set(c, "order repo", err)
			}
			currentOrderID = newOrderID
		}
	}
	//創建transaction
	tx := p.transactionRepo.CreateTransaction()
	//存入收據
	item := response.LatestReceiptInfo[0]
	if item.ExpiresDate == nil {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	quantity, err := strconv.Atoi(item.Quantity)
	if err != nil {
		return p.errHandler.Set(c, "Atoi error", err)
	}
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               currentOrderID,
		PaymentType:           int(global.ApplePaymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: item.OriginalTransactionID,
		TransactionID:         item.TransactionID,
		ProductID:             item.ProductID,
		Quantity:              quantity,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//獲取訂閱項目資訊
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(item.ProductID)
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "subscribe plan error", err)
	}
	//綁定訂單與會員關係
	var subscribeStatus = global.ValidSubscribeStatus
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		subscribeStatus = global.NoneSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          order.UserID,
		OrderID:         currentOrderID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       item.PurchaseDate.Format("2006-01-02 15:04:05"),
		ExpiresDate:     item.ExpiresDate.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "member repo", err)
	}
	//更新會員類型
	var userType = global.SubscribeUserType
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		userType = global.NormalUserType
	}
	ut := int(userType)
	if err := p.userRepo.UpdateUserByUID(tx, order.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	//更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, currentOrderID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func (p *payment) handleSubscribeTradeForGoogle(c *gin.Context, uid int64, order *dto.SubscribeOrder, receiptData string, response *dto.GoogleReceiptResponse) errcode.Error {
	//初始化當前訂單id
	currentOrderID := order.ID
	//查詢用戶訂閱資訊
	subscribeInfo, err := p.subscribeInfo.FindSubscribeInfo(uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return p.errHandler.Set(c, "user subscribe repo", err)
	}
	if subscribeInfo != nil {
		if subscribeInfo.OrderID != nil {
			currentOrderID = *subscribeInfo.OrderID
		}
	}
	//驗證訂單初始交易id
	receipts, err := p.receiptRepo.FindReceiptsByOrderID(currentOrderID, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return p.errHandler.Set(c, "receipt repo", err)
	}
	if len(receipts) > 0 {
		return nil
	}
	//創建transaction
	tx := p.transactionRepo.CreateTransaction()
	//存入收據
	if err != nil {
		return p.errHandler.Set(c, "Atoi error", err)
	}
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               currentOrderID,
		PaymentType:           int(global.ApplePaymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: response.OrderID,
		TransactionID:         response.OrderID,
		ProductID:             order.SubscribePlan.ProductID,
		Quantity:              1,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//獲取訂閱項目資訊
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(order.SubscribePlan.ProductID)
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "subscribe plan error", err)
	}
	//綁定訂單與會員關係
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          order.UserID,
		OrderID:         currentOrderID,
		SubscribePlanID: subscribePlan.ID,
		Status:          global.ValidSubscribeStatus,
		StartDate:       time.Now().Format("2006-01-02 15:04:05"),
		ExpiresDate:     time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "member repo", err)
	}
	//更新會員類型
	ut := int(global.SubscribeUserType)
	if err := p.userRepo.UpdateUserByUID(tx, order.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	//更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, currentOrderID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func parserIAPDate(unix int64) *time.Time {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(unix, 0).Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return nil
	}
	return &date
}

func parserCourseOrder(data *model.Order) *dto.CourseOrder {
	if data == nil {
		return nil
	}
	order := dto.CourseOrder{
		ID:          data.ID,
		UserID:      data.UserID,
		Quantity:    data.Quantity,
		OrderType:   data.OrderType,
		OrderStatus: data.OrderStatus,
		CreateAt:    data.CreateAt,
		UpdateAt:    data.UpdateAt,
	}
	if data.OrderCourse != nil {
		if data.OrderCourse.SaleItem != nil {
			order.SaleItem = &dto.SaleItem{
				ID:   data.OrderCourse.SaleItem.ID,
				Type: data.OrderCourse.SaleItem.Type,
			}
			if data.OrderCourse.SaleItem.ProductLabel != nil {
				order.SaleItem.Name = data.OrderCourse.SaleItem.Name
				order.SaleItem.ProductID = data.OrderCourse.SaleItem.ProductLabel.ProductID
				order.SaleItem.Twd = data.OrderCourse.SaleItem.ProductLabel.Twd
			}
		}
		if data.OrderCourse.Course != nil {
			order.Course = &dto.CourseProductSummary{
				ID:           data.OrderCourse.Course.ID,
				SaleType:     data.OrderCourse.Course.SaleType,
				CourseStatus: data.OrderCourse.Course.CourseStatus,
				Category:     data.OrderCourse.Course.Category,
				ScheduleType: data.OrderCourse.Course.ScheduleType,
				Name:         data.OrderCourse.Course.Name,
				Cover:        data.OrderCourse.Course.Cover,
				Level:        data.OrderCourse.Course.Level,
				PlanCount:    data.OrderCourse.Course.PlanCount,
				WorkoutCount: data.OrderCourse.Course.WorkoutCount,
			}
			order.Course.Review.ScoreTotal = data.OrderCourse.Course.Review.ScoreTotal
			order.Course.Review.Amount = data.OrderCourse.Course.Review.Amount
			if data.OrderCourse.Course.Trainer != nil {
				order.Course.Trainer = &dto.TrainerSummary{
					UserID:   data.OrderCourse.Course.Trainer.UserID,
					Nickname: data.OrderCourse.Course.Trainer.Nickname,
					Avatar:   data.OrderCourse.Course.Trainer.Avatar,
					Skill:    data.OrderCourse.Course.Trainer.Skill,
				}
			}
			if data.OrderCourse.Course.Sale != nil {
				order.Course.Sale = &dto.SaleItem{
					ID:   data.OrderCourse.Course.Sale.ID,
					Type: data.OrderCourse.Course.Sale.Type,
					Name: data.OrderCourse.Course.Sale.Name,
				}
				if data.OrderCourse.Course.Sale.ProductLabel != nil {
					order.Course.Sale.ProductID = data.OrderCourse.Course.Sale.ProductLabel.ProductID
					order.Course.Sale.Twd = data.OrderCourse.Course.Sale.ProductLabel.Twd
				}
			}
		}
	}
	return &order
}

func parserSubscribeOrder(data *model.Order) *dto.SubscribeOrder {
	order := dto.SubscribeOrder{
		ID:          data.ID,
		UserID:      data.UserID,
		Quantity:    data.Quantity,
		OrderType:   data.OrderType,
		OrderStatus: data.OrderStatus,
		CreateAt:    data.CreateAt,
		UpdateAt:    data.UpdateAt,
	}
	if data.OrderSubscribe != nil {
		if data.OrderSubscribe.SubscribePlan != nil {
			order.SubscribePlan = &dto.SubscribePlan{
				ID:     data.OrderSubscribe.SubscribePlan.ID,
				Period: data.OrderSubscribe.SubscribePlan.Period,
			}
			if data.OrderSubscribe.SubscribePlan.ProductLabel != nil {
				order.SubscribePlan.ProductID = data.OrderSubscribe.SubscribePlan.ProductLabel.ProductID
				order.SubscribePlan.Name = data.OrderSubscribe.SubscribePlan.ProductLabel.Name
				order.SubscribePlan.Twd = data.OrderSubscribe.SubscribePlan.ProductLabel.Twd
			}
		}
	}
	return &order
}

func randRange(min int64, max int64) int64 {
	if min > max || min < 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}
