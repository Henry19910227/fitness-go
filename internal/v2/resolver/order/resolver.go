package order

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iab"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iap"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	iabModel "github.com/Henry19910227/fitness-go/internal/v2/model/iab"
	iapModel "github.com/Henry19910227/fitness-go/internal/v2/model/iap"
	orderModel "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderCourseModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	orderSubscribePlanModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	purchaseLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/purchase_log"
	receiptModel "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	subscribeLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
	subscribePlanModel "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	userModel "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	courseAssetModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/purchase_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/service/subscribe_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type resolver struct {
	orderService              order.Service
	courseService             course.Service
	orderCourserService       order_course.Service
	courseAssetService        user_course_asset.Service
	subscribeInfoService      user_subscribe_info.Service
	orderSubscribePlanService order_subscribe_plan.Service
	receiptService            receipt.Service
	purchaseLogService        purchase_log.Service
	subscribePlanService      subscribe_plan.Service
	userService               user.Service
	subscribeLogService       subscribe_log.Service
	iapTool                   iap.Tool
	iabTool                   iab.Tool
}

func New(orderService order.Service, courseService course.Service,
	orderCourserService order_course.Service, courseAssetService user_course_asset.Service,
	subscribeInfoService user_subscribe_info.Service, orderSubscribePlanService order_subscribe_plan.Service,
	receiptService receipt.Service, purchaseLogService purchase_log.Service,
	subscribePlanService subscribe_plan.Service, userService user.Service,
	subscribeLogService subscribe_log.Service, iapTool iap.Tool, iabTool iab.Tool) Resolver {
	return &resolver{orderService: orderService, courseService: courseService,
		orderCourserService: orderCourserService, courseAssetService: courseAssetService,
		subscribeInfoService: subscribeInfoService, orderSubscribePlanService: orderSubscribePlanService,
		receiptService: receiptService, purchaseLogService: purchaseLogService,
		subscribePlanService: subscribePlanService, userService: userService,
		subscribeLogService: subscribeLogService, iapTool: iapTool, iabTool: iabTool}
}

func (r *resolver) APICreateCourseOrder(tx *gorm.DB, input *orderModel.APICreateCourseOrderInput) (output orderModel.APICreateCourseOrderOutput) {
	// 檢查是此課表是否已購買
	findAssetsInput := courseAssetModel.ListInput{}
	findAssetsInput.UserID = util.PointerInt64(input.UserID)
	findAssetsInput.CourseID = util.PointerInt64(input.Body.CourseID)
	findAssetsInput.Available = util.PointerInt(1)
	assetOutputs, _, err := r.courseAssetService.List(&findAssetsInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(assetOutputs) > 0 {
		output.Set(code.BadRequest, "此課表已被購買，無法再創建訂單")
		return output
	}
	// 檢查此課表狀態
	courseFindInput := courseModel.FindInput{}
	courseFindInput.ID = util.PointerInt64(input.Body.CourseID)
	courseOutput, err := r.courseService.Find(&courseFindInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.CourseStatus, 0) != courseModel.Sale {
		output.Set(code.BadRequest, "必須為銷售狀態的課表才可被加入訂單")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypeFree && util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypeCharge {
		output.Set(code.BadRequest, "商品必須為免費課表或付費課表類型才可創建此訂單")
		return output
	}
	if (util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypeFree) && (courseOutput.SaleID == nil) {
		output.Set(code.BadRequest, "付費課表必須有 sale id")
		return output
	}
	// 檢查是否有尚未付款的相同商品訂單，如有則直接回傳
	orderListInput := orderModel.ListInput{}
	orderListInput.UserID = util.PointerInt64(input.UserID)
	orderListInput.CourseID = util.PointerInt64(input.Body.CourseID)
	orderListInput.OrderStatus = util.PointerInt(orderModel.Pending)
	orderListInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
	}
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) > 0 {
		data := orderModel.APICreateCourseOrderData{}
		if err := util.Parser(orderOutputs[0], &data); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	// 產出訂單流水號
	orderID := time.Now().Format("20060102150405") + strconv.Itoa(int(util.RandRange(100000, 999999)))
	// 創建課表訂單
	defer tx.Rollback()
	table := orderModel.Table{}
	table.ID = util.PointerString(orderID)
	table.UserID = util.PointerInt64(input.UserID)
	table.Quantity = util.PointerInt(1)
	table.Type = util.PointerInt(orderModel.BuyCourse)
	table.OrderStatus = util.PointerInt(orderModel.Pending)
	id, err := r.orderService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建訂單與課表關聯
	orderCourseTable := orderCourseModel.Table{}
	orderCourseTable.OrderID = util.PointerString(id)
	orderCourseTable.SaleItemID = courseOutput.SaleID
	orderCourseTable.CourseID = util.PointerInt64(input.Body.CourseID)
	if err := r.orderCourserService.Tx(tx).Create(&orderCourseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// find order return
	findOrdersInput := orderModel.FindInput{}
	findOrdersInput.ID = util.PointerString(id)
	findOrdersInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
	}
	orderOutput, err := r.orderService.Find(&findOrdersInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := orderModel.APICreateCourseOrderData{}
	if err := util.Parser(orderOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APICreateSubscribeOrder(tx *gorm.DB, input *orderModel.APICreateSubscribeOrderInput) (output orderModel.APICreateSubscribeOrderOutput) {
	// 檢查目前是否已訂閱
	subscribeListInput := subscribeInfoModel.ListInput{}
	subscribeListInput.UserID = util.PointerInt64(input.UserID)
	subscribeListInput.Page = 1
	subscribeListInput.Size = 1
	subscribeListOutput, _, err := r.subscribeInfoService.List(&subscribeListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(subscribeListOutput) > 0 {
		if util.OnNilJustReturnInt(subscribeListOutput[0].Status, 0) == 1 {
			output.Set(code.BadRequest, "目前已是訂閱會員")
			return output
		}
	}
	// 檢查是否有創建過訂閱訂單，有則修改訂單訂閱項目並回傳(一個用戶最多只會產生一個訂閱的訂單)
	orderListInput := orderModel.ListInput{}
	orderListInput.UserID = util.PointerInt64(input.UserID)
	orderListInput.Type = util.PointerInt(orderModel.Subscribe)
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderListInput.Page = 1
	orderListInput.Size = 1
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) > 0 {
		// 修改訂單訂閱項目
		orderSubscribePlanTable := orderSubscribePlanModel.Table{}
		orderSubscribePlanTable.OrderID = orderOutputs[0].ID
		orderSubscribePlanTable.SubscribePlanID = util.PointerInt64(input.Body.SubscribePlanID)
		if err := r.orderSubscribePlanService.Update(&orderSubscribePlanTable); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 查找該訂單
		findOrderInput := orderModel.FindInput{}
		findOrderInput.ID = orderOutputs[0].ID
		findOrderInput.Preloads = []*preloadModel.Preload{
			{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
		}
		orderOutput, err := r.orderService.Find(&findOrderInput)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// Parser output
		data := orderModel.APICreateSubscribeOrderData{}
		if err := util.Parser(orderOutput, &data); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	// 產出訂單流水號
	orderID := time.Now().Format("20060102150405") + strconv.Itoa(int(util.RandRange(100000, 999999)))
	// 創建訂閱訂單
	defer tx.Rollback()
	table := orderModel.Table{}
	table.ID = util.PointerString(orderID)
	table.UserID = util.PointerInt64(input.UserID)
	table.Quantity = util.PointerInt(1)
	table.Type = util.PointerInt(orderModel.Subscribe)
	table.OrderStatus = util.PointerInt(orderModel.Pending)
	id, err := r.orderService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建訂單與訂閱項目關聯
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = util.PointerString(id)
	orderSubscribePlanTable.SubscribePlanID = util.PointerInt64(input.Body.SubscribePlanID)
	if err := r.orderSubscribePlanService.Tx(tx).Create(&orderSubscribePlanTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// find order return
	findOrderInput := orderModel.FindInput{}
	findOrderInput.ID = util.PointerString(id)
	findOrderInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	}
	orderOutput, err := r.orderService.Find(&findOrderInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := orderModel.APICreateSubscribeOrderData{}
	if err := util.Parser(orderOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIVerifyAppleReceipt(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIVerifyAppleReceiptInput) (output orderModel.APIVerifyAppleReceiptOutput) {
	findInput := orderModel.FindInput{}
	findInput.ID = util.PointerString(input.Body.OrderID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	}
	orderOutput, err := r.orderService.Find(&findInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(orderOutput.UserID, 0) != input.UserID {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+"無效的訂單，此訂單非該用戶創建")
		output.Set(code.BadRequest, "無效的訂單，此訂單非該用戶創建")
		return output
	}
	response, err := r.iapTool.VerifyAppleReceiptAPI(input.Body.ReceiptData)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//驗證收據結果
	if response.Status != 0 {
		//更新訂單狀態
		orderTable := orderModel.Table{}
		orderTable.OrderStatus = util.PointerInt(orderModel.Error)
		_ = r.orderService.Update(&orderTable)
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+"收據驗證錯誤 - "+strconv.Itoa(response.Status))
		output.Set(code.BadRequest, "收據驗證錯誤 - "+strconv.Itoa(response.Status))
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) == orderModel.BuyCourse {
		if err := r.handleBuyCourseTradeForApple(tx, orderOutput, response); err != nil {
			logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) == orderModel.Subscribe {
		if err := r.handleSubscribeTradeForApple(tx, orderOutput, response); err != nil {
			logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		return output
	}
	logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+"訂單類型錯誤")
	output.Set(code.BadRequest, "訂單類型錯誤")
	return output
}

func (r *resolver) APIVerifyGoogleReceipt(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIVerifyGoogleReceiptInput) (output orderModel.APIVerifyGoogleReceiptOutput) {
	findInput := orderModel.FindInput{}
	findInput.ID = util.PointerString(input.Body.OrderID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	}
	orderOutput, err := r.orderService.Find(&findInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(orderOutput.UserID, 0) != input.UserID {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+"無效的訂單，此訂單非該用戶創建")
		output.Set(code.BadRequest, "無效的訂單，此訂單非該用戶創建")
		return output
	}
	//產出 auth token
	oauthToken, err := r.iabTool.GenerateGoogleOAuth2Token(time.Hour)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取API Token
	token, err := r.iabTool.APIGetGooglePlayToken(oauthToken)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) == orderModel.BuyCourse {
		response, err := r.iabTool.APIGetProducts(input.Body.ProductID, input.Body.ReceiptData, token)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
		}
		if err := r.handleBuyCourseTradeForGoogle(tx, orderOutput, response); err != nil {
			logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) == orderModel.Subscribe {
		response, err := r.iabTool.APIGetSubscription(input.Body.ProductID, input.Body.ReceiptData, token)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
		}
		if err := r.handleSubscribeTradeForGoogle(tx, orderOutput, response); err != nil {
			logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		return output
	}
	logger.Shared().Error(ctx, "APIVerifyGoogleReceipt："+"訂單類型錯誤")
	output.Set(code.BadRequest, "訂單類型錯誤")
	return output
}

func (r *resolver) APIAppStoreNotification(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIAppStoreNotificationInput) (output orderModel.APIAppStoreNotificationOutput) {
	//解析字串
	response := dto.NewIAPNotificationResponse(strings.Split(input.Body.SignedPayload, ".")[1])
	if response == nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+"SignedPayload 解析錯誤(response 為 null)")
		output.Set(code.BadRequest, "SignedPayload 解析錯誤")
		return output
	}
	if response.Data == nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+"SignedPayload 解析錯誤(response.Data 為 null)")
		output.Set(code.BadRequest, "SignedPayload 解析錯誤")
		return output
	}
	if response.Data.SignedTransactionInfo == nil || response.Data.SignedRenewalInfo == nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+"SignedPayload 解析錯誤(response.Data.SignedTransactionInfo 或 response.Data.SignedRenewalInfo 為 null)")
		output.Set(code.BadRequest, "SignedPayload 解析錯誤")
		return output
	}
	//存取 subscribe log
	subscribeLogType := subscribeLogModel.GetTypeByIAPType(response.NotificationType, response.Subtype)
	subscribeLogTable := subscribeLogModel.Table{}
	subscribeLogTable.OriginalTransactionID = util.PointerString(response.Data.SignedTransactionInfo.OriginalTransactionId)
	subscribeLogTable.TransactionID = util.PointerString(response.Data.SignedTransactionInfo.TransactionId)
	subscribeLogTable.PurchaseDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
	subscribeLogTable.ExpiresDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
	subscribeLogTable.Type = util.PointerString(subscribeLogType)
	subscribeLogTable.Msg = util.PointerString(fmt.Sprintf("%s %s", response.NotificationType, response.Subtype))
	_, err := r.subscribeLogService.CreateOrUpdate(&subscribeLogTable)
	if err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 獲取訂閱項目資訊
	findPlanInput := subscribePlanModel.FindInput{}
	findPlanInput.ProductID = util.PointerString(response.Data.SignedTransactionInfo.ProductId)
	subscribePlanOutput, err := r.subscribePlanService.Find(&findPlanInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢 OriginalTransactionId 關聯訂單
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(response.Data.SignedTransactionInfo.OriginalTransactionId)
	orderListInput.Size = 1
	orderListInput.Page = 1
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) == 0 {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+"查無關聯訂單")
		output.Set(code.BadRequest, "查無關聯訂單")
		return output
	}
	defer tx.Rollback()
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = orderOutputs[0].ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAP)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.OriginalTransactionID = util.PointerString(response.Data.SignedTransactionInfo.OriginalTransactionId)
	receiptTable.TransactionID = util.PointerString(response.Data.SignedTransactionInfo.TransactionId)
	receiptTable.ProductID = util.PointerString(response.Data.SignedTransactionInfo.ProductId)
	receiptTable.Quantity = util.PointerInt(int(response.Data.SignedTransactionInfo.Quantity))
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.修改訂單訂閱項目(升級or降級狀態)
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = orderOutputs[0].ID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3.更新用戶訂閱狀態
	var subscribeStatus = subscribeInfoModel.ValidSubscribe
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		subscribeStatus = subscribeInfoModel.NoneSubscribe
	}
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = orderOutputs[0].UserID
	subscribeInfoTable.OrderID = orderOutputs[0].ID
	subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4.更新用戶類型
	var userType = userModel.Subscribe
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		userType = userModel.Normal
	}
	userTable := userModel.Table{}
	userTable.ID = orderOutputs[0].UserID
	userTable.UserType = util.PointerInt(userType)
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = orderOutputs[0].ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGooglePlayNotification(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIGooglePlayNotificationInput) (output orderModel.APIGooglePlayNotificationOutput) {
	//解析字串
	notificationResp := iabModel.NewIABSubscribeNotificationResponse(input.Body.Message.Data)
	if notificationResp == nil {
		logger.Shared().Error(ctx, "iab notification decode error")
		output.Set(code.BadRequest, "iab notification decode error")
		return output
	}
	if notificationResp.SubscriptionNotification == nil {
		logger.Shared().Error(ctx, "iab notification decode error")
		output.Set(code.BadRequest, "iab notification decode error")
		return output
	}
	//產出 auth token
	oauthToken, err := r.iabTool.GenerateGoogleOAuth2Token(time.Hour)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取API Token
	token, err := r.iabTool.APIGetGooglePlayToken(oauthToken)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//驗證收據
	response, err := r.iabTool.APIGetSubscription(notificationResp.SubscriptionNotification.SubscriptionId, notificationResp.SubscriptionNotification.PurchaseToken, token)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取 originalTransactionID 與 transactionID
	originalTransactionID := response.OrderId
	transactionID := response.OrderId
	transactionIDs := strings.Split(response.OrderId, "..")
	if len(transactionIDs) > 1 {
		originalTransactionID = transactionIDs[0]
	}
	//存取 subscribe log
	startTimeMillis, err := strconv.ParseInt(response.StartTimeMillis, 10, 64)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	expiryTimeMillis, err := strconv.ParseInt(response.ExpiryTimeMillis, 10, 64)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	subscribeLogTable := subscribeLogModel.Table{}
	subscribeLogTable.OriginalTransactionID = util.PointerString(originalTransactionID)
	subscribeLogTable.TransactionID = util.PointerString(transactionID)
	subscribeLogTable.PurchaseDate = util.PointerString(util.UnixToTime(startTimeMillis).Format("2006-01-02 15:04:05"))
	subscribeLogTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis).Format("2006-01-02 15:04:05"))
	subscribeLogTable.Type = util.PointerString(subscribeLogModel.GetTypeByIABType(notificationResp.SubscriptionNotification.NotificationType))
	subscribeLogTable.Msg = util.PointerString(subscribeLogModel.GetMsgByIABType(notificationResp.SubscriptionNotification.NotificationType))
	_, err = r.subscribeLogService.CreateOrUpdate(&subscribeLogTable)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 獲取訂閱項目資訊
	findPlanInput := subscribePlanModel.FindInput{}
	findPlanInput.ProductID = util.PointerString(notificationResp.SubscriptionNotification.SubscriptionId)
	subscribePlanOutput, err := r.subscribePlanService.Find(&findPlanInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢 OriginalTransactionId 關聯訂單
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(originalTransactionID)
	orderListInput.Size = 1
	orderListInput.Page = 1
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) == 0 {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+"查無關聯訂單")
		output.Set(code.BadRequest, "查無關聯訂單")
		return output
	}
	defer tx.Rollback()
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = orderOutputs[0].ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAB)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.OriginalTransactionID = util.PointerString(originalTransactionID)
	receiptTable.TransactionID = util.PointerString(transactionID)
	receiptTable.ProductID = util.PointerString(notificationResp.SubscriptionNotification.SubscriptionId)
	receiptTable.Quantity = util.PointerInt(1)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.修改訂單訂閱項目(升級or降級狀態)
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = orderOutputs[0].ID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3.更新用戶訂閱狀態
	subscribeLogType := subscribeLogModel.GetTypeByIABType(notificationResp.SubscriptionNotification.NotificationType)
	var subscribeStatus = subscribeInfoModel.ValidSubscribe
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		subscribeStatus = subscribeInfoModel.NoneSubscribe
	}
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = orderOutputs[0].UserID
	subscribeInfoTable.OrderID = orderOutputs[0].ID
	subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(startTimeMillis).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis).Format("2006-01-02 15:04:05"))
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4.更新用戶類型
	var userType = userModel.Subscribe
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		userType = userModel.Normal
	}
	userTable := userModel.Table{}
	userTable.ID = orderOutputs[0].UserID
	userTable.UserType = util.PointerInt(userType)
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = orderOutputs[0].ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIVerifyAppleSubscribe(input *orderModel.APIVerifyAppleSubscribeInput) (output orderModel.APIVerifyAppleSubscribeOutput) {
	// 驗證是否是原先訂閱用戶
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(input.Body.OriginalTransactionID)
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderListInput.Size = 1
	orderListInput.Page = 1
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) > 0 {
		if util.OnNilJustReturnInt64(orderOutputs[0].UserID, 0) != input.UserID {
			output.Set(code.BadRequest, "無法訂閱，此 Apple ID 已綁定其他帳戶")
			return output
		}
	}
	output.Set(code.Success, "可進行訂閱")
	return output
}

func (r *resolver) APIGetCMSOrders(input *orderModel.APIGetCMSOrdersInput) (output orderModel.APIGetCMSOrdersOutput) {
	// parser input
	param := orderModel.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// get list
	datas, page, err := r.orderService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := orderModel.APIGetCMSOrdersData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) handleBuyCourseTradeForApple(tx *gorm.DB, order *orderModel.Output, response *iapModel.IAPVerifyReceiptResponse) error {
	//驗證收據格式
	if len(response.Receipt.InApp) == 0 {
		return errors.New("無效的收據(無InApp參數)")
	}
	item := response.Receipt.InApp[0]
	if util.OnNilJustReturnString(order.OrderCourse.SaleItem.ProductLabel.ProductID, "") != item.ProductID {
		return errors.New("無效的收據(與訂單 ProductID 不匹配)")
	}
	//驗證收據是否已被使用
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	receiptOutputs, _, err := r.receiptService.List(&receiptListInput)
	if err != nil {
		return err
	}
	if len(receiptOutputs) > 0 {
		return errors.New("此收據已有支付記錄")
	}
	defer tx.Rollback()
	//創建收據
	quantity, _ := strconv.Atoi(item.Quantity)
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = order.ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAP)
	receiptTable.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	receiptTable.TransactionID = util.PointerString(item.TransactionID)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.ProductID = util.PointerString(item.ProductID)
	receiptTable.Quantity = util.PointerInt(quantity)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		return err
	}
	//創建購買紀錄
	logTable := purchaseLogModel.Table{}
	logTable.UserID = order.UserID
	logTable.OrderID = order.ID
	logTable.Type = util.PointerInt(purchaseLogModel.Buy)
	_, err = r.purchaseLogService.Tx(tx).Create(&logTable)
	if err != nil {
		return err
	}
	//創建購買資源
	assetTable := courseAssetModel.Table{}
	assetTable.UserID = order.UserID
	assetTable.CourseID = order.OrderCourse.CourseID
	assetTable.Available = util.PointerInt(1)
	_, err = r.courseAssetService.Tx(tx).Create(&assetTable)
	if err != nil {
		return err
	}
	//更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = order.ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (r *resolver) handleBuyCourseTradeForGoogle(tx *gorm.DB, order *orderModel.Output, response *iabModel.IABProductAPIResponse) error {
	//驗證收據格式
	if response.PurchaseState != 0 {
		return errors.New("尚未購買")
	}
	//驗證收據是否已被使用
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.OriginalTransactionID = util.PointerString(response.OrderId)
	receiptOutputs, _, err := r.receiptService.List(&receiptListInput)
	if err != nil {
		return err
	}
	if len(receiptOutputs) > 0 {
		return errors.New("此收據已有支付記錄")
	}
	defer tx.Rollback()
	//創建收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = order.ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAB)
	receiptTable.OriginalTransactionID = util.PointerString(response.OrderId)
	receiptTable.TransactionID = util.PointerString(response.OrderId)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.ProductID = order.OrderCourse.SaleItem.ProductLabel.ProductID
	receiptTable.Quantity = util.PointerInt(1)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		return err
	}
	//創建購買紀錄
	logTable := purchaseLogModel.Table{}
	logTable.UserID = order.UserID
	logTable.OrderID = order.ID
	logTable.Type = util.PointerInt(purchaseLogModel.Buy)
	_, err = r.purchaseLogService.Tx(tx).Create(&logTable)
	if err != nil {
		return err
	}
	//創建購買資源
	assetTable := courseAssetModel.Table{}
	assetTable.UserID = order.UserID
	assetTable.CourseID = order.OrderCourse.CourseID
	assetTable.Available = util.PointerInt(1)
	_, err = r.courseAssetService.Tx(tx).Create(&assetTable)
	if err != nil {
		return err
	}
	//更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = order.ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (r *resolver) handleSubscribeTradeForApple(tx *gorm.DB, order *orderModel.Output, response *iapModel.IAPVerifyReceiptResponse) error {
	// 驗證收據格式
	if len(response.LatestReceiptInfo) == 0 || len(response.PendingRenewalInfo) == 0 {
		return errors.New("無效的收據(無LatestReceiptInfo或PendingRenewalInfo參數)")
	}
	item := response.LatestReceiptInfo[0]
	if response.LatestReceiptInfo[0].ExpiresDate == nil {
		return errors.New("無效的收據(無ExpiresDate參數)")
	}
	// 驗證是否是原先訂閱用戶
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderListInput.Size = 1
	orderListInput.Page = 1
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		return err
	}
	if len(orderOutputs) > 0 {
		if util.OnNilJustReturnInt64(orderOutputs[0].UserID, 0) != util.OnNilJustReturnInt64(order.UserID, 0) {
			return errors.New("驗證失敗(與原先訂閱用戶不符)")
		}
	}
	// 獲取訂閱項目資訊
	findPlanInput := subscribePlanModel.FindInput{}
	findPlanInput.ProductID = util.PointerString(item.ProductID)
	subscribePlanOutput, err := r.subscribePlanService.Find(&findPlanInput)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// 1.儲存收據
	quantity, _ := strconv.Atoi(item.Quantity)
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = order.ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAP)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	receiptTable.TransactionID = util.PointerString(item.TransactionID)
	receiptTable.ProductID = util.PointerString(item.ProductID)
	receiptTable.Quantity = util.PointerInt(quantity)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		return err
	}
	// 2.修改訂單訂閱項目(升級or降級狀態)
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = order.ID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		return err
	}
	// 3.更新用戶訂閱狀態
	var subscribeStatus = subscribeInfoModel.ValidSubscribe
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		subscribeStatus = subscribeInfoModel.NoneSubscribe
	}
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = order.UserID
	subscribeInfoTable.OrderID = order.ID
	subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
	subscribeInfoTable.StartDate = util.PointerString(item.PurchaseDate.Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(item.ExpiresDate.Format("2006-01-02 15:04:05"))
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		return err
	}
	// 4.更新用戶類型
	var userType = userModel.Subscribe
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		userType = userModel.Normal
	}
	userTable := userModel.Table{}
	userTable.ID = order.UserID
	userTable.UserType = util.PointerInt(userType)
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		return err
	}
	// 5.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = order.ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (r *resolver) handleSubscribeTradeForGoogle(tx *gorm.DB, order *orderModel.Output, response *iabModel.IABSubscriptionAPIResponse) error {
	// 驗證是否是原先訂閱用戶
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(response.OrderId)
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderListInput.Size = 1
	orderListInput.Page = 1
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		return err
	}
	if len(orderOutputs) > 0 {
		if util.OnNilJustReturnInt64(orderOutputs[0].UserID, 0) != util.OnNilJustReturnInt64(order.UserID, 0) {
			return errors.New("驗證失敗(與原先訂閱用戶不符)")
		}
	}
	// 獲取訂閱項目資訊
	findPlanInput := subscribePlanModel.FindInput{}
	findPlanInput.ProductID = order.OrderCourse.SaleItem.ProductLabel.ProductID
	subscribePlanOutput, err := r.subscribePlanService.Find(&findPlanInput)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// 回傳的訂單如遇到 GPA.3331-2251-2804-48618..4
	// OriginalTransactionID = 只留下'..'前的訂單編號 GPA.3331-2251-2804-48618
	// TransactionID = 完整的訂單編號 GPA.3331-2251-2804-48618
	originalTransactionID := response.OrderId
	transactionID := response.OrderId
	transactionIDs := strings.Split(response.OrderId, "..")
	if len(transactionIDs) > 1 {
		originalTransactionID = transactionIDs[0]
	}
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = order.ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAB)
	receiptTable.OriginalTransactionID = util.PointerString(originalTransactionID)
	receiptTable.TransactionID = util.PointerString(transactionID)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.ProductID = order.OrderCourse.SaleItem.ProductLabel.ProductID
	receiptTable.Quantity = util.PointerInt(1)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		return err
	}
	// 2.修改訂單訂閱項目(升級or降級狀態)
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = order.ID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		return err
	}
	// 3.更新用戶訂閱狀態
	var subscribeStatus = subscribeInfoModel.ValidSubscribe
	if response.PaymentState == 1 || response.PaymentState == 2 {
		subscribeStatus = subscribeInfoModel.NoneSubscribe
	}
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = order.UserID
	subscribeInfoTable.OrderID = order.ID
	subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
	startTimeMillis, err := strconv.ParseInt(response.StartTimeMillis, 10, 64)
	if err != nil {
		return nil
	}
	expiryTimeMillis, err := strconv.ParseInt(response.ExpiryTimeMillis, 10, 64)
	if err != nil {
		return nil
	}
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(startTimeMillis).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis).Format("2006-01-02 15:04:05"))
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		return err
	}
	// 4.更新用戶類型
	var userType = userModel.Subscribe
	if response.PaymentState == 1 || response.PaymentState == 2 {
		userType = userModel.Normal
	}
	userTable := userModel.Table{}
	userTable.ID = order.UserID
	userTable.UserType = util.PointerInt(userType)
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		return err
	}
	// 5.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = order.ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		return err
	}
	tx.Commit()
	return nil
}
