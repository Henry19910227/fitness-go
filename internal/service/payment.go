package service

import (
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
	"strconv"
	"strings"
	"time"
)

type payment struct {
	Base
	userRepo            repository.User
	orderRepo           repository.Order
	orderSPRepo         repository.OrderSubscribePlan
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
	iabHandler          handler.IAB
	reqTool             tool.HttpRequest
	jwtTool             tool.JWT
	errHandler          errcode.Handler
}

func NewPayment(userRepo repository.User, orderRepo repository.Order,
	orderSPRepo repository.OrderSubscribePlan,
	saleRepo repository.Sale, subscribePlanRepo repository.SubscribePlan,
	courseRepo repository.Course, receiptRepo repository.Receipt,
	purchaseRepo repository.UserCourseAsset, subscribeLogRepo repository.SubscribeLog,
	purchaseLogRepo repository.PurchaseLog, memberRepo repository.UserSubscribeInfo,
	transactionRepo repository.Transaction, iapHandler handler.IAP, iabHandler handler.IAB, reqTool tool.HttpRequest,
	jwtTool tool.JWT, errHandler errcode.Handler) Payment {
	return &payment{userRepo: userRepo, orderRepo: orderRepo, orderSPRepo: orderSPRepo,
		saleRepo: saleRepo, subscribePlanRepo: subscribePlanRepo,
		courseRepo: courseRepo, receiptRepo: receiptRepo,
		userCourseAssetRepo: purchaseRepo, subscribeLogRepo: subscribeLogRepo, purchaseLogRepo: purchaseLogRepo,
		subscribeInfo: memberRepo, transactionRepo: transactionRepo,
		reqTool: reqTool, jwtTool: jwtTool, iapHandler: iapHandler, iabHandler: iabHandler, errHandler: errHandler}
}

func (p *payment) GetAppleStoreApiAccessToken(c *gin.Context) (string, errcode.Error) {
	accessToken, err := p.iapHandler.GetAppleStoreAPIAccessToken()
	if err != nil {
		return "", p.errHandler.Set(c, "iap handler", err)
	}
	return accessToken, nil
}

func (p *payment) GetGooglePlayApiAccessToken(c *gin.Context) (string, errcode.Error) {
	accessToken, err := p.iabHandler.GetGooglePlayApiAccessToken()
	if err != nil {
		return "", p.errHandler.Set(c, "iab handler", err)
	}
	return accessToken, nil
}

func (p *payment) GetAppStoreAPISubscriptions(c *gin.Context, originalTransactionID string) (*dto.IAPSubscribeAPIResponse, errcode.Error) {
	result, err := p.iapHandler.GetSubscribeAPI(originalTransactionID)
	if err != nil {
		return nil, p.errHandler.Set(c, "iap handler", err)
	}
	return result, nil
}

func (p *payment) GetAppStoreAPIHistory(c *gin.Context, originalTransactionID string) (*dto.IAPHistoryAPIResponse, errcode.Error) {
	result, err := p.iapHandler.GetHistoryAPI(originalTransactionID)
	if err != nil {
		return nil, p.errHandler.Set(c, "iap handler", err)
	}
	return result, nil
}

func (p *payment) GetGooglePlayAPIProduct(c *gin.Context, productID string, purchaseToken string) (*dto.IABProductAPIResponse, errcode.Error) {
	result, err := p.iabHandler.GetProductsAPI(productID, purchaseToken)
	if err != nil {
		return nil, p.errHandler.Set(c, "get products api error", err)
	}
	return result, nil
}

func (p *payment) GetGooglePlayAPISubscription(c *gin.Context, productID string, purchaseToken string) (*dto.IABSubscriptionAPIResponse, errcode.Error) {
	result, err := p.iabHandler.GetSubscriptionAPI(productID, purchaseToken)
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
	//驗證該用戶訂閱狀態
	subscribeStatus, e := p.CheckUserSubscribeStatus(c, uid)
	if e != nil {
		return nil, e
	}
	if subscribeStatus == global.ValidSubscribeStatus {
		return nil, p.errHandler.Custom(8999, errors.New("目前已是訂閱會員"))
	}
	//查詢當前訂閱項目
	subscribePlanData, err := p.subscribePlanRepo.FinsSubscribePlanByID(subscribePlanID)
	if err != nil {
		return nil, p.errHandler.Set(c, "subscribe plan repo", err)
	}
	subscribePlan := dto.NewSubscribePlan(subscribePlanData)
	//查找該用戶是否創建過訂閱訂單
	var paymentOrderType = global.SubscribeOrderType
	orders, err := p.orderRepo.FindOrders(uid, &model.FindOrdersParam{
		PaymentOrderType: &paymentOrderType,
	}, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	if len(orders) > 0 {
		order := dto.NewSubscribeOrder(orders[0])
		order.SubscribePlan = &subscribePlan
		return &order, nil
	}
	//創建新的訂單
	orderID, err := p.orderRepo.CreateSubscribeOrder(uid)
	if err != nil {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	orderData, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	order := dto.NewSubscribeOrder(orderData)
	order.SubscribePlan = &subscribePlan
	return &order, nil
}

func (p *payment) CheckUserSubscribeStatus(c *gin.Context, uid int64) (global.SubscribeStatus, errcode.Error) {
	//查詢該用戶最新Apple收據
	appleReceipts, err := p.receiptRepo.FindReceiptsByPaymentType(uid, global.ApplePaymentType, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return global.NoneSubscribeStatus, p.errHandler.Set(c, "receipt repo", err)
	}
	//驗證該apple收據訂閱狀態
	if len(appleReceipts) > 0 {
		response, err := p.iapHandler.GetSubscribeAPI(appleReceipts[0].OriginalTransactionID)
		if err != nil {
			return global.NoneSubscribeStatus, nil
		}
		if response != nil {
			if len(response.Data) > 0 {
				if len(response.Data[0].LastTransactions) > 0 {
					status := response.Data[0].LastTransactions[0].Status
					if status == 1 || status == 3 || status == 4 || status == 5 { // 當前訂閱尚未過期
						return global.ValidSubscribeStatus, nil
					}
				}
			}
		}
	}
	//查詢該用戶最新google收據
	googleReceipts, err := p.receiptRepo.FindReceiptsByPaymentType(uid, global.GooglePaymentType, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		return global.NoneSubscribeStatus, p.errHandler.Set(c, "receipt repo", err)
	}
	//驗證該google收據訂閱狀態
	if len(googleReceipts) > 0 {
		response, err := p.iabHandler.GetSubscriptionAPI(googleReceipts[0].ProductID, googleReceipts[0].ReceiptToken)
		if err != nil {
			return global.NoneSubscribeStatus, nil
		}
		if response.PaymentState == 1 || response.PaymentState == 2 {
			return global.ValidSubscribeStatus, nil
		}
	}
	return global.NoneSubscribeStatus, nil
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
		return p.errHandler.Custom(8999, errors.New("無效的訂單(此訂單非本人)"))
	}
	//驗證收據API
	response, err := p.iapHandler.VerifyAppleReceiptAPI(receiptData)
	if err != nil {
		return p.errHandler.Set(c, "iap handler", err)
	}
	//驗證收據結果
	if response.Status != 0 {
		//更新訂單狀態
		_ = p.orderRepo.UpdateOrderStatus(nil, order.ID, global.ErrorOrderStatus)
		return p.errHandler.Custom(8999, errors.New("收據驗證錯誤"))
	}
	//處理課表購買訂單
	if order.OrderType == int(global.BuyCourseOrderType) {
		if err := p.handleBuyCourseTrade(c, parserCourseOrder(order), receiptData, response); err != nil {
			return err
		}
	}
	//處理會員訂閱訂單
	if order.OrderType == int(global.SubscribeOrderType) {
		if err := p.handleSubscribeTradeForApple(c, uid, parserSubscribeOrder(order), receiptData, response); err != nil {
			return err
		}
	}
	return nil
}

func (p *payment) VerifyGoogleReceipt(c *gin.Context, uid int64, orderID string, productID string, receiptData string) errcode.Error {
	//取得訂單資訊
	order, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	if order.UserID != uid {
		return p.errHandler.Custom(8999, errors.New("無效的收據(此訂單非本人)"))
	}
	//處理課表購買訂單
	if order.OrderType == int(global.BuyCourseOrderType) {
		if err := p.handleBuyCourseTradeForGoogle(c, parserCourseOrder(order), productID, receiptData); err != nil {
			return err
		}
	}
	//處理會員訂閱訂單
	if order.OrderType == int(global.SubscribeOrderType) {
		if err := p.handleSubscribeTradeForGoogle(c, uid, parserSubscribeOrder(order), productID, receiptData); err != nil {
			return err
		}
	}
	return nil
}

func (p *payment) HandleAppStoreNotification(c *gin.Context, base64PayloadString string) errcode.Error {
	//解析字串
	response := dto.NewIAPNotificationResponse(strings.Split(base64PayloadString, ".")[1])
	if response == nil {
		return p.errHandler.Set(c, "iap notification", errors.New("iap notification decode error"))
	}
	subscribeLogType := p.iapHandler.ParserIAPNotificationType(response.NotificationType, response.Subtype)
	//存取 subscribe log
	if response.Data == nil {
		return p.errHandler.Custom(8999, errors.New("error parser notification"))
	}
	if response.Data.SignedTransactionInfo == nil || response.Data.SignedRenewalInfo == nil {
		return p.errHandler.Custom(8999, errors.New("error parser notification"))
	}
	_, err := p.subscribeLogRepo.SaveSubscribeLog(nil, &model.CreateSubscribeLogParam{
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
	// 查詢當前訂閱項目
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(response.Data.SignedTransactionInfo.ProductId)
	if err != nil {
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//以 OriginalTransactionId 查詢訂單
	order, err := p.orderRepo.FindOrderByOriginalTransactionID(response.Data.SignedTransactionInfo.OriginalTransactionId)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	tx := p.transactionRepo.CreateTransaction()
	//1.存取收據
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
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
	//2.修改訂單訂閱項目(升級or降級狀態)
	if err := p.orderSPRepo.SaveOrderSubscribePlan(tx, order.ID, subscribePlan.ID); err != nil {
		tx.Rollback()
		return nil
	}
	//3.修改用戶訂閱狀態
	var subscribeStatus = global.ValidSubscribeStatus
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		subscribeStatus = global.NoneSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          order.UserID,
		OrderID:         order.ID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       parserIAPDate(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"),
		ExpiresDate:     parserIAPDate(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "subscribe info repo", err)
	}
	//4.更新會員類型
	var userType = global.SubscribeUserType
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		userType = global.NormalUserType
	}
	ut := int(userType)
	if err := p.userRepo.UpdateUserByUID(tx, order.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	//5.更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, order.ID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func (p *payment) HandleGooglePlayNotification(c *gin.Context, base64PayloadString string) errcode.Error {
	//解析字串
	notificationResp := dto.NewIABSubscribeNotificationResponse(base64PayloadString)
	if notificationResp == nil {
		return p.errHandler.Set(c, "iab notification", errors.New("iab notification decode error"))
	}
	if notificationResp.SubscriptionNotification == nil {
		return p.errHandler.Set(c, "iab notification", errors.New("iab notification decode error"))
	}
	//google play 驗證收據 API
	resp, err := p.iabHandler.GetSubscriptionAPI(notificationResp.SubscriptionNotification.SubscriptionId, notificationResp.SubscriptionNotification.PurchaseToken)
	if err != nil {
		return p.errHandler.Set(c, "iab notification", err)
	}
	//獲取 originalTransactionID 與 transactionID
	originalTransactionID := resp.OrderId
	transactionID := resp.OrderId
	transactionIDs := strings.Split(resp.OrderId, "..")
	if len(transactionIDs) > 1 {
		originalTransactionID = transactionIDs[0]
	}
	//存取 subscribe log
	subscribeLogType := p.iabHandler.ParserIABNotificationType(notificationResp.SubscriptionNotification.NotificationType)
	_, err = p.subscribeLogRepo.SaveSubscribeLog(nil, &model.CreateSubscribeLogParam{
		OriginalTransactionID: originalTransactionID,
		TransactionID:         transactionID,
		PurchaseDate:          parserDate(resp.StartTimeMillis).Format("2006-01-02 15:04:05"),
		ExpiresDate:           parserDate(resp.ExpiryTimeMillis).Format("2006-01-02 15:04:05"),
		Type:                  string(subscribeLogType),
		Msg:                   p.iabHandler.ParserIABNotificationMsg(notificationResp.SubscriptionNotification.NotificationType),
	})
	if err != nil {
		return p.errHandler.Set(c, "subscribe log repo", err)
	}
	// 以SubscriptionId查詢當前訂閱項目
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(notificationResp.SubscriptionNotification.SubscriptionId)
	if err != nil {
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//以 OriginalTransactionId 查詢訂單
	order, err := p.orderRepo.FindOrderByOriginalTransactionID(originalTransactionID)
	if err != nil {
		return p.errHandler.Set(c, "order repo", err)
	}
	tx := p.transactionRepo.CreateTransaction()
	defer p.transactionRepo.FinishTransaction(tx)
	//1.存取收據
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
		PaymentType:           int(global.GooglePaymentType),
		ReceiptToken:          "",
		OriginalTransactionID: originalTransactionID,
		TransactionID:         transactionID,
		ProductID:             notificationResp.SubscriptionNotification.SubscriptionId,
		Quantity:              1,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//2.修改訂單訂閱項目(升級or降級狀態)
	if err := p.orderSPRepo.SaveOrderSubscribePlan(tx, order.ID, subscribePlan.ID); err != nil {
		tx.Rollback()
		return nil
	}
	//3.修改用戶訂閱狀態
	var subscribeStatus = global.ValidSubscribeStatus
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		subscribeStatus = global.NoneSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          order.UserID,
		OrderID:         order.ID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       parserDate(resp.StartTimeMillis).Format("2006-01-02 15:04:05"),
		ExpiresDate:     parserDate(resp.ExpiryTimeMillis).Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "subscribe info repo", err)
	}
	//4.更新會員類型
	var userType = global.SubscribeUserType
	if subscribeLogType == global.Expired || subscribeLogType == global.Refund {
		userType = global.NormalUserType
	}
	ut := int(userType)
	if err := p.userRepo.UpdateUserByUID(tx, order.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	//5.更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, order.ID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	return nil
}

func (p *payment) newSubscribeOrder(userID int64, subscribePlanID int64) (*dto.SubscribeOrder, error) {
	subscribePlanData, err := p.subscribePlanRepo.FinsSubscribePlanByID(subscribePlanID)
	if err != nil {
		return nil, err
	}
	//創建新的訂單
	orderID, err := p.orderRepo.CreateSubscribeOrder(userID)
	if err != nil {
		return nil, err
	}
	// 查找剛創建的訂單
	orderData, err := p.orderRepo.FindOrder(orderID)
	if err != nil {
		return nil, err
	}
	order := dto.NewSubscribeOrder(orderData)
	subscribePlan := dto.NewSubscribePlan(subscribePlanData)
	order.SubscribePlan = &subscribePlan
	return parserSubscribeOrder(orderData), nil
}

func (p *payment) handleFreeCourseTrade(c *gin.Context, order *dto.CourseOrder) errcode.Error {
	if err := p.handleCourseOrderTrade(order, global.NonePaymentType, "",
		"", "", 1, ""); err != nil {
		return p.errHandler.Set(c, "trade error", err)
	}
	return nil
}

func (p *payment) handleBuyCourseTrade(c *gin.Context, order *dto.CourseOrder, receiptData string, response *dto.IAPVerifyReceiptResponse) errcode.Error {
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

func (p *payment) handleBuyCourseTradeForGoogle(c *gin.Context, order *dto.CourseOrder, purchaseProductID string, receiptData string) errcode.Error {
	//驗證產品ID
	if order.SaleItem == nil {
		return p.errHandler.Custom(8999, errors.New("訂單缺少價格資訊"))
	}
	if order.SaleItem.ProductID != purchaseProductID {
		return p.errHandler.Custom(8999, errors.New("訂單與產品ID不一致"))
	}
	//Google Play API 驗證收據
	response, err := p.iabHandler.GetProductsAPI(purchaseProductID, receiptData)
	if err != nil {
		return p.errHandler.Custom(8999, errors.New("收據驗證失敗"))
	}
	//驗證購買狀態
	if response.PurchaseState != 0 {
		return p.errHandler.Custom(8999, errors.New("尚未購買"))
	}
	//處理課表最終交易
	if err := p.handleCourseOrderTrade(order, global.GooglePaymentType, response.OrderId,
		response.OrderId, order.SaleItem.ProductID, 1, receiptData); err != nil {
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

func (p *payment) handleSubscribeTradeForApple(c *gin.Context, uid int64, order *dto.SubscribeOrder, receiptData string, response *dto.IAPVerifyReceiptResponse) errcode.Error {
	//驗證收據格式
	if len(response.LatestReceiptInfo) == 0 || len(response.PendingRenewalInfo) == 0 {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	item := response.LatestReceiptInfo[0]
	if response.LatestReceiptInfo[0].ExpiresDate == nil {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	//驗證訂單合法性
	if order.UserID != uid {
		return p.errHandler.Custom(8999, errors.New("此收據與該用戶不符"))
	}
	//獲取訂閱項目資訊
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(item.ProductID)
	if err != nil {
		return p.errHandler.Set(c, "subscribe plan error", err)
	}
	tx := p.transactionRepo.CreateTransaction()
	//1.存入收據
	quantity, err := strconv.Atoi(item.Quantity)
	if err != nil {
		return p.errHandler.Set(c, "Atoi error", err)
	}
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
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
	//2.修改訂單訂閱項目(升級or降級狀態)
	if err := p.orderSPRepo.SaveOrderSubscribePlan(tx, order.ID, subscribePlan.ID); err != nil {
		tx.Rollback()
		return nil
	}
	//3.更新用戶訂閱狀態
	var subscribeStatus = global.ValidSubscribeStatus
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		subscribeStatus = global.NoneSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          uid,
		OrderID:         order.ID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       item.PurchaseDate.Format("2006-01-02 15:04:05"),
		ExpiresDate:     item.ExpiresDate.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "member repo", err)
	}
	//4.更新用戶類型
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
	//5.更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, order.ID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func (p *payment) handleSubscribeTradeForGoogle(c *gin.Context, uid int64, order *dto.SubscribeOrder, purchaseProductID string, receiptData string) errcode.Error {
	//獲取訂閱項目資訊
	subscribePlan, err := p.subscribePlanRepo.FinsSubscribePlanByProductID(purchaseProductID)
	if err != nil {
		return p.errHandler.Set(c, "subscribe plan error", err)
	}
	//Google Play API 驗證收據
	response, err := p.iabHandler.GetSubscriptionAPI(purchaseProductID, receiptData)
	if err != nil {
		return p.errHandler.Custom(8999, errors.New("收據驗證失敗"))
	}
	//創建transaction
	tx := p.transactionRepo.CreateTransaction()
	defer p.transactionRepo.FinishTransaction(tx)
	// 回傳的訂單如遇到 GPA.3331-2251-2804-48618..4
	// OriginalTransactionID = 只留下'..'前的訂單編號 GPA.3331-2251-2804-48618
	// TransactionID = 完整的訂單編號 GPA.3331-2251-2804-48618
	originalTransactionID := response.OrderId
	transactionID := response.OrderId
	transactionIDs := strings.Split(response.OrderId, "..")
	if len(transactionIDs) > 1 {
		originalTransactionID = transactionIDs[0]
	}
	//1.存入收據
	_, err = p.receiptRepo.SaveReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
		PaymentType:           int(global.GooglePaymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: originalTransactionID,
		TransactionID:         transactionID,
		ProductID:             subscribePlan.ProductLabel.ProductID,
		Quantity:              1,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//2.修改訂單訂閱項目(升級or降級狀態)
	if err := p.orderSPRepo.SaveOrderSubscribePlan(tx, order.ID, subscribePlan.ID); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order sp repo", err)
	}
	//3.更新用戶訂閱狀態
	var subscribeStatus = global.NoneSubscribeStatus
	if response.PaymentState == 1 || response.PaymentState == 2 {
		subscribeStatus = global.ValidSubscribeStatus
	}
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:          uid,
		OrderID:         order.ID,
		SubscribePlanID: subscribePlan.ID,
		Status:          subscribeStatus,
		StartDate:       parserDate(response.StartTimeMillis).Format("2006-01-02 15:04:05"),
		ExpiresDate:     parserDate(response.ExpiryTimeMillis).Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "member repo", err)
	}
	//4.更新用戶類型
	var userType = global.NormalUserType
	if response.PaymentState == 1 || response.PaymentState == 2 {
		userType = global.SubscribeUserType
	}
	ut := int(userType)
	if err := p.userRepo.UpdateUserByUID(tx, order.UserID, &model.UpdateUserParam{
		UserType: &ut,
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "user repo", err)
	}
	//5.更新訂單狀態
	if err := p.orderRepo.UpdateOrderStatus(tx, order.ID, global.SuccessOrderStatus); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
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

func parserDate(unixMS string) *time.Time {
	msTime, err := strconv.ParseInt(unixMS, 10, 64)
	if err != nil {
		return nil
	}
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(msTime/1000, 0).Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return nil
	}
	return &date
}
