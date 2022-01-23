package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type payment struct {
	Base
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
	reqTool             tool.HttpRequest
	jwtTool             tool.JWT
	errHandler          errcode.Handler
}

func NewPayment(orderRepo repository.Order, saleRepo repository.Sale, subscribePlanRepo repository.SubscribePlan,
	courseRepo repository.Course, receiptRepo repository.Receipt,
	purchaseRepo repository.UserCourseAsset, subscribeLogRepo repository.SubscribeLog,
	purchaseLogRepo  repository.PurchaseLog, memberRepo repository.UserSubscribeInfo,
	transactionRepo  repository.Transaction, reqTool tool.HttpRequest,
	jwtTool tool.JWT, errHandler errcode.Handler) Payment {
	return &payment{orderRepo: orderRepo, saleRepo: saleRepo, subscribePlanRepo: subscribePlanRepo,
		courseRepo: courseRepo, receiptRepo: receiptRepo,
		userCourseAssetRepo: purchaseRepo, subscribeLogRepo: subscribeLogRepo, purchaseLogRepo: purchaseLogRepo,
		subscribeInfo: memberRepo, transactionRepo: transactionRepo,
		reqTool: reqTool, jwtTool: jwtTool, errHandler: errHandler}
}

func (p *payment) Test(c *gin.Context) (string, errcode.Error) {
	result, err := p.jwtTool.GenerateAppleToken()
	if err != nil {
		return "", p.errHandler.Set(c, "", err)
	}
	return result, nil
}

func (p *payment) CreateCourseOrder(c *gin.Context, uid int64, courseID int64) (*dto.CourseOrder, errcode.Error) {
	//檢查是此課表是否已購買
	courseAsset, err := p.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
		UserID: uid,
		CourseID: courseID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
		return nil, p.errHandler.Set(c, "order repo", err)
	}
	if courseAsset != nil {
		if courseAsset.Available == 1 {
			return nil, p.errHandler.Custom(8999, errors.New("此課表已被購買，無法再創建訂單"))
		}
	}
	//檢查是否有尚未付款的相同訂單
	orderData, err := p.orderRepo.FindOrderByUserIDAndCourseID(uid, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
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
	if course.SaleID == nil {
		return nil, p.errHandler.Custom(8999, errors.New("付費訂單的課表必須有 sale id"))
	}
	// 創建訂單
	orderID, err := p.orderRepo.CreateCourseOrder(&model.CreateOrderParam{
		UserID: uid,
		SaleItemID: *course.SaleID,
		CourseID: courseID,
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

func (p *payment) CreateSubscribeOrder(c *gin.Context, uid int64, period global.PeriodType) (*dto.SubscribeOrder, errcode.Error) {
	//驗證當前訂閱狀態
	subscribeInfo, err := p.subscribeInfo.FindSubscribeInfo(uid)
	if err != nil {
		return nil, p.errHandler.Set(c, "subscribe info repo", err)
	}
	if global.SubscribeStatus(subscribeInfo.Status) == global.ValidSubscribeStatus {
		return nil, p.errHandler.Custom(8999, errors.New("目前已經是訂閱會員"))
	}
	// 查詢訂閱方案
	plans, err := p.subscribePlanRepo.FindSubscribePlansByPeriod(period)
	if err != nil {
		return nil, p.errHandler.Set(c, "sale item repo", err)
	}
	if len(plans) == 0 {
		return nil, p.errHandler.Custom(8999, errors.New("無此訂閱方案"))
	}
	plan := plans[0]
	// 創建訂單
	orderID, err := p.orderRepo.CreateSubscribeOrder(&model.CreateSubscribeOrderParam{
		UserID:     uid,
		SubscribePlanID: plan.ID,
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
	//驗證訂單狀態
	if global.OrderStatus(order.OrderStatus) != global.PendingOrderStatus {
		return p.errHandler.Custom(8999, errors.New("此訂單已失效"))
	}
	//apple server 正式區驗證收據
	param := map[string]interface{}{
		"receipt-data": receiptData,
		"password": "b3e50e11316943969754106ed24c6a3a",
		"exclude-old-transactions": 1,
	}
	result, err := p.reqTool.SendPostRequestWithJsonBody("https://buy.itunes.apple.com/verifyReceipt", param)
	if err != nil {
		return p.errHandler.Set(c, "req tool", err)
	}
	var response dto.AppleReceiptResponse
	if err := parserAppleReceipt(result, &response); err != nil {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	//apple server 測試區驗證收據
	if response.Status == 21007 {
		result, err := p.reqTool.SendPostRequestWithJsonBody("https://sandbox.itunes.apple.com/verifyReceipt", param)
		if err != nil {
			return p.errHandler.Set(c, "req tool", err)
		}
		if err := parserAppleReceipt(result, &response); err != nil {
			return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
		}
	}
	//驗證收據結果
	if response.Status != 0 {
		//更新訂單狀態
		_ = p.orderRepo.UpdateOrder(nil, order.ID, &model.UpdateOrderParam{
			OrderStatus: int(global.ErrorOrderStatus),
		})
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
		if err := p.handleSubscribeTrade(c, parserSubscribeOrder(order), receiptData, &response); err != nil {
			return err
		}
	}
	return nil
}

func (p *payment) HandleAppStoreNotification(c *gin.Context, base64PayloadString string) (*dto.IAPNotificationResponse, errcode.Error) {
	response, err := parserIAPNotificationResponse(base64PayloadString)
	if err != nil {
		p.errHandler.Set(c, "iap parser", err)
	}
	fmt.Printf("NotificationType: %v \n", response.NotificationType)
	fmt.Printf("Subtype: %v \n", response.Subtype)

	fmt.Printf("Data.Environment: %v \n", response.Data.Environment)


	fmt.Printf("SignedRenewalInfo.AutoRenewProductId: %v \n", response.Data.SignedRenewalInfo.AutoRenewProductId)
	fmt.Printf("SignedRenewalInfo.AutoRenewStatus: %v \n", response.Data.SignedRenewalInfo.AutoRenewStatus)
	fmt.Printf("SignedRenewalInfo.ExpirationIntent: %v \n", response.Data.SignedRenewalInfo.ExpirationIntent)
	fmt.Printf("SignedRenewalInfo.GracePeriodExpiresDate: %v \n", parserIAPDate(response.Data.SignedRenewalInfo.GracePeriodExpiresDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedRenewalInfo.IsInBillingRetryPeriod: %v \n", response.Data.SignedRenewalInfo.IsInBillingRetryPeriod)
	fmt.Printf("SignedRenewalInfo.OfferIdentifier: %v \n", response.Data.SignedRenewalInfo.OfferIdentifier)
	fmt.Printf("SignedRenewalInfo.OfferType: %v \n", response.Data.SignedRenewalInfo.OfferType)
	fmt.Printf("SignedRenewalInfo.OriginalTransactionId: %v \n", response.Data.SignedRenewalInfo.OriginalTransactionId)
	fmt.Printf("SignedRenewalInfo.PriceIncreaseStatus: %v \n", response.Data.SignedRenewalInfo.PriceIncreaseStatus)
	fmt.Printf("SignedRenewalInfo.ProductId: %v \n", response.Data.SignedRenewalInfo.ProductId)
	fmt.Printf("SignedRenewalInfo.SignedDate: %v \n", parserIAPDate(response.Data.SignedRenewalInfo.SignedDate / 1000).Format("2006-01-02 15:04:05"))


	fmt.Printf("SignedTransactionInfo.AppAccountToken: %v \n", response.Data.SignedTransactionInfo.AppAccountToken)
	fmt.Printf("SignedTransactionInfo.BundleId: %v \n", response.Data.SignedTransactionInfo.BundleId)
	fmt.Printf("SignedTransactionInfo.ExpiresDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.InAppOwnershipType: %v \n", response.Data.SignedTransactionInfo.InAppOwnershipType)
	fmt.Printf("SignedTransactionInfo.IsUpgraded: %v \n", response.Data.SignedTransactionInfo.IsUpgraded)
	fmt.Printf("SignedTransactionInfo.OfferIdentifier: %v \n", response.Data.SignedTransactionInfo.OfferIdentifier)
	fmt.Printf("SignedTransactionInfo.OfferType: %v \n", response.Data.SignedTransactionInfo.OfferType)
	fmt.Printf("SignedTransactionInfo.OriginalPurchaseDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.OriginalPurchaseDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.OriginalTransactionId: %v \n", response.Data.SignedTransactionInfo.OriginalTransactionId)
	fmt.Printf("SignedTransactionInfo.ProductId: %v \n", response.Data.SignedTransactionInfo.ProductId)
	fmt.Printf("SignedTransactionInfo.PurchaseDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.Quantity: %v \n", response.Data.SignedTransactionInfo.Quantity)
	fmt.Printf("SignedTransactionInfo.RevocationDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.RevocationDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.RevocationReason: %v \n", response.Data.SignedTransactionInfo.RevocationReason)
	fmt.Printf("SignedTransactionInfo.SignedDate: %v \n", parserIAPDate(response.Data.SignedTransactionInfo.SignedDate / 1000).Format("2006-01-02 15:04:05"))
	fmt.Printf("SignedTransactionInfo.SubscriptionGroupIdentifier: %v \n", response.Data.SignedTransactionInfo.SubscriptionGroupIdentifier)
	fmt.Printf("SignedTransactionInfo.TransactionId: %v \n", response.Data.SignedTransactionInfo.TransactionId)
	fmt.Printf("SignedTransactionInfo.Type: %v \n", response.Data.SignedTransactionInfo.Type)
	fmt.Printf("SignedTransactionInfo.WebOrderLineItemId: %v \n", response.Data.SignedTransactionInfo.WebOrderLineItemId)

	return response, nil
}

func (p *payment) handleFreeCourseTrade(c *gin.Context, order *dto.CourseOrder) errcode.Error {
	if err := p.handleCourseOrderTrade(c, order, global.NonePaymentType, "",
		"", 1, ""); err != nil {
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
	if err := p.handleCourseOrderTrade(c, order, global.ApplePaymentType, item.OriginalTransactionID,
		item.TransactionID, quantity, receiptData); err != nil {
		return p.errHandler.Set(c, "trade error", err)
	}
	return nil
}

func (p *payment) handleCourseOrderTrade(c *gin.Context, order *dto.CourseOrder, paymentType global.PaymentType,
	originalTransactionID string, transactionID string, quantity int, receiptData string) error {
	//創建transaction
	tx := p.transactionRepo.CreateTransaction()
	//存入收據
	_, err := p.receiptRepo.CreateReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
		PaymentType:           int(paymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: originalTransactionID,
		TransactionID:         transactionID,
		Quantity:              quantity,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	//存入購買Log
	_, err = p.purchaseLogRepo.CreatePurchaseLog(tx, &model.CreatePurchaseLogParam{
		UserID: order.UserID,
		OrderID: order.ID,
		Type: global.BuyPurchaseLogType,
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
	if err := p.orderRepo.UpdateOrder(tx, order.ID, &model.UpdateOrderParam{
		OrderStatus: int(global.SuccessOrderStatus),
	}); err != nil {
		tx.Rollback()
		return err
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func (p *payment) handleSubscribeTrade(c *gin.Context, order *dto.SubscribeOrder, receiptData string, response *dto.AppleReceiptResponse) errcode.Error {
	if len(response.LatestReceiptInfo) == 0 {
		return p.errHandler.Custom(8999, errors.New("收據格式錯誤"))
	}
	//驗證產品id
	if order.SubscribePlan.ProductID != response.LatestReceiptInfo[0].ProductID {
		return p.errHandler.Custom(8999, errors.New("無效的收據(產品ID不匹配)"))
	}
	//驗證當前訂閱狀態
	member, err := p.subscribeInfo.FindSubscribeInfo(order.UserID)
	if err != nil {
		return p.errHandler.Set(c, "member repo", err)
	}
	if global.SubscribeStatus(member.Status) == global.ValidSubscribeStatus {
		return p.errHandler.Custom(8999, errors.New("目前已經是訂閱會員"))
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
	_, err = p.receiptRepo.CreateReceipt(tx, &model.CreateReceiptParam{
		OrderID:               order.ID,
		PaymentType:           int(global.ApplePaymentType),
		ReceiptToken:          receiptData,
		OriginalTransactionID: item.OriginalTransactionID,
		TransactionID:         item.TransactionID,
		Quantity:              quantity,
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "receipt repo", err)
	}
	//存入訂閱紀錄
	_, err = p.subscribeLogRepo.CreateSubscribeLog(tx, &model.CreateSubscribeLogParam{
		UserID:  order.UserID,
		TransactionID: item.TransactionID,
		PurchaseDate: item.PurchaseDate.Format("2006-01-02 15:04:05"),
		ExpiresDate: item.ExpiresDate.Format("2006-01-02 15:04:05"),
		Type:    string(global.NormalSubscribeLogType),
		Msg:     "訂閱成功!",
	})
	if err != nil {
		tx.Rollback()
		return  p.errHandler.Set(c, "subscribe log repo", err)
	}
	//更新會員資料
	_, err = p.subscribeInfo.SaveSubscribeInfo(tx, &model.SaveUserSubscribeInfoParam{
		UserID:       order.UserID,
		Status: global.ValidSubscribeStatus,
		StartDate: item.PurchaseDate.Format("2006-01-02 15:04:05"),
		ExpiresDate: item.ExpiresDate.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "member repo", err)
	}
	//更新訂單狀態
	if err := p.orderRepo.UpdateOrder(tx, order.ID, &model.UpdateOrderParam{
		OrderStatus: int(global.SuccessOrderStatus),
	}); err != nil {
		tx.Rollback()
		return p.errHandler.Set(c, "order repo", err)
	}
	//結束transaction
	p.transactionRepo.FinishTransaction(tx)
	return nil
}

func parserIAPNotificationResponse(base64String string) (*dto.IAPNotificationResponse, error) {
	payloadDict, err := decodeBase64StringToMap(strings.Split(base64String, ".")[1])
	if err != nil {
		return nil, err
	}
	payloadDataDict, ok := payloadDict["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedRenewalInfoString, ok := payloadDataDict["signedRenewalInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedTransactionInfoString, ok := payloadDataDict["signedTransactionInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	renewalInfoDict, err := decodeBase64StringToMap(strings.Split(base64SignedRenewalInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	transactionInfo, err := decodeBase64StringToMap(strings.Split(base64SignedTransactionInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	// parser dto
	var response dto.IAPNotificationResponse
	if err := mapstructure.Decode(payloadDict, &response); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(payloadDataDict, &response.Data); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(renewalInfoDict, &response.Data.SignedRenewalInfo); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(transactionInfo, &response.Data.SignedTransactionInfo); err != nil {
		return nil, err
	}
	return &response, nil
}

func decodeBase64StringToMap(base64String string) (map[string]interface{}, error) {
	payloadString, err := base64.RawURLEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(payloadString, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}

func parserAppleReceipt(dict map[string]interface{}, receipt *dto.AppleReceiptResponse) error {
	if err := mapstructure.Decode(dict, &receipt); err != nil {
		return err
	}
	for _, item := range receipt.Receipt.InApp {
		parserAppleReceiptInfo(item)
	}
	for _, item := range receipt.LatestReceiptInfo {
		parserAppleReceiptInfo(item)
	}
	return nil
}

func parserAppleReceiptInfo(receiptInfo *dto.ReceiptInfo) {
	originalPurchaseDate, err := parserAppleReceiptDate(receiptInfo.OriginalPurchaseDateMS)
	if err == nil {
		receiptInfo.OriginalPurchaseDate = originalPurchaseDate
	}
	purchaseDate, err := parserAppleReceiptDate(receiptInfo.PurchaseDateMS)
	if err == nil {
		receiptInfo.PurchaseDate = purchaseDate
	}
	expiresDate, err := parserAppleReceiptDate(receiptInfo.ExpiresDateMS)
	if err == nil {
		receiptInfo.ExpiresDate = expiresDate
	}
}

func parserAppleReceiptDate(unixMS string) (*time.Time, error) {
	msTime, err := strconv.ParseInt(unixMS, 10, 64)
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(msTime / 1000, 0).Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return nil, err
	}
	return &date, nil
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
	if data == nil { return nil }
	order := dto.CourseOrder{
		ID: data.ID,
		UserID: data.UserID,
		Quantity: data.Quantity,
		OrderType: data.OrderType,
		OrderStatus: data.OrderStatus,
		CreateAt: data.CreateAt,
		UpdateAt: data.UpdateAt,
	}
	if data.OrderCourse != nil {
		if data.OrderCourse.SaleItem != nil {
			order.SaleItem = &dto.SaleItem{
				ID: data.OrderCourse.SaleItem.ID,
				Type: data.OrderCourse.SaleItem.Type,
			}
			if data.OrderCourse.SaleItem.ProductLabel != nil {
				order.SaleItem.Name = data.OrderCourse.SaleItem.Name
				order.SaleItem.ProductID = data.OrderCourse.SaleItem.ProductLabel.ProductID
				order.SaleItem.Twd = data.OrderCourse.SaleItem.ProductLabel.Twd
			}
		}
		if data.OrderCourse.Course != nil {
			order.Course = &dto.CourseProductItem{
				ID: data.OrderCourse.Course.ID,
				SaleType: data.OrderCourse.Course.SaleType,
				Name: data.OrderCourse.Course.Name,
				Cover: data.OrderCourse.Course.Cover,
			}
		}
	}
	return &order
}

func parserSubscribeOrder(data *model.Order) *dto.SubscribeOrder {
	order := dto.SubscribeOrder{
		ID: data.ID,
		UserID: data.UserID,
		Quantity: data.Quantity,
		OrderType: data.OrderType,
		OrderStatus: data.OrderStatus,
		CreateAt: data.CreateAt,
		UpdateAt: data.UpdateAt,
	}
	if data.OrderSubscribe != nil {
		if data.OrderSubscribe.SubscribePlan != nil {
			order.SubscribePlan = &dto.SubscribePlan{
				ID: data.OrderSubscribe.SubscribePlan.ID,
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