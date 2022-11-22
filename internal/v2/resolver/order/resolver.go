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
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderModel "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_create_subscribe_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_order_redeem"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_charge_receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order/api_upload_apple_subscribe_receipt"
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
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/purchase_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
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
	saleItemService           sale_item.Service
	iapTool                   iap.Tool
	iabTool                   iab.Tool
}

func New(orderService order.Service, courseService course.Service,
	orderCourserService order_course.Service, courseAssetService user_course_asset.Service,
	subscribeInfoService user_subscribe_info.Service, orderSubscribePlanService order_subscribe_plan.Service,
	receiptService receipt.Service, purchaseLogService purchase_log.Service,
	subscribePlanService subscribe_plan.Service, userService user.Service,
	subscribeLogService subscribe_log.Service, saleItemService sale_item.Service,
	iapTool iap.Tool, iabTool iab.Tool) Resolver {
	return &resolver{orderService: orderService, courseService: courseService,
		orderCourserService: orderCourserService, courseAssetService: courseAssetService,
		subscribeInfoService: subscribeInfoService, orderSubscribePlanService: orderSubscribePlanService,
		receiptService: receiptService, purchaseLogService: purchaseLogService,
		subscribePlanService: subscribePlanService, userService: userService,
		subscribeLogService: subscribeLogService, saleItemService: saleItemService,
		iapTool: iapTool, iabTool: iabTool}
}

func (r *resolver) APICreateCourseOrder(tx *gorm.DB, input *orderModel.APICreateCourseOrderInput) (output orderModel.APICreateCourseOrderOutput) {
	defer tx.Rollback()
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

func (r *resolver) APICreateSubscribeOrder(tx *gorm.DB, input *api_create_subscribe_order.Input) (output api_create_subscribe_order.Output) {
	defer tx.Rollback()
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
	//// 檢查是否有創建過訂閱訂單，有則修改訂單訂閱項目並回傳(一個用戶最多只會產生一個訂閱的訂單)
	//orderListInput := orderModel.ListInput{}
	//orderListInput.UserID = util.PointerInt64(input.UserID)
	//orderListInput.Type = util.PointerInt(orderModel.Subscribe)
	//orderListInput.OrderField = "create_at"
	//orderListInput.OrderType = order_by.DESC
	//orderListInput.Page = 1
	//orderListInput.Size = 1
	//orderOutputs, _, err := r.orderService.List(&orderListInput)
	//if err != nil {
	//	output.Set(code.BadRequest, err.Error())
	//	return output
	//}
	//if len(orderOutputs) > 0 {
	//	// 修改訂單訂閱項目
	//	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	//	orderSubscribePlanTable.OrderID = orderOutputs[0].ID
	//	orderSubscribePlanTable.SubscribePlanID = util.PointerInt64(input.Body.SubscribePlanID)
	//	if err := r.orderSubscribePlanService.Update(&orderSubscribePlanTable); err != nil {
	//		output.Set(code.BadRequest, err.Error())
	//		return output
	//	}
	//	// 查找該訂單
	//	findOrderInput := orderModel.FindInput{}
	//	findOrderInput.ID = orderOutputs[0].ID
	//	findOrderInput.Preloads = []*preloadModel.Preload{
	//		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	//	}
	//	orderOutput, err := r.orderService.Find(&findOrderInput)
	//	if err != nil {
	//		output.Set(code.BadRequest, err.Error())
	//		return output
	//	}
	//	// Parser output
	//	data := api_create_subscribe_order.Data{}
	//	if err := util.Parser(orderOutput, &data); err != nil {
	//		output.Set(code.BadRequest, err.Error())
	//		return output
	//	}
	//	output.Set(code.Success, "success")
	//	output.Data = &data
	//	return output
	//}
	// 產出訂單流水號
	orderID := time.Now().Format("20060102150405") + strconv.Itoa(int(util.RandRange(100000, 999999)))
	// 創建訂閱訂單
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
	orderSubscribePlanTable.Status = util.PointerInt(0)
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
	data := api_create_subscribe_order.Data{}
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

func (r *resolver) APIUploadAppleSubscribeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_apple_subscribe_receipt.Input) (output api_upload_apple_subscribe_receipt.Output) {
	// 驗證收據 api
	response, err := r.iapTool.VerifyAppleReceiptAPI(input.Body.ReceiptData)
	if err != nil {
		logger.Shared().Error(ctx, "APIUploadAppleSubscribeReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證收據結果
	if response.Status != 0 {
		logger.Shared().Error(ctx, "APIUploadAppleSubscribeReceipt："+"收據驗證錯誤 - "+strconv.Itoa(response.Status))
		output.Set(code.BadRequest, "收據驗證錯誤 - "+strconv.Itoa(response.Status))
		return output
	}
	// 驗證收據格式
	if len(response.Receipt.InApp) == 0 {
		output.Set(code.BadRequest, "無效的收據(無InApp參數)")
		return output
	}
	if len(response.LatestReceiptInfo) == 0 || len(response.PendingRenewalInfo) == 0 {
		output.Set(code.BadRequest, "無效的收據(無LatestReceiptInfo或PendingRenewalInfo參數)")
		return output
	}
	if response.LatestReceiptInfo[0].ExpiresDate == nil {
		output.Set(code.BadRequest, "無效的收據(無ExpiresDate參數)")
		return output
	}
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		output.Set(code.BadRequest, "此訂閱收據已過期")
		return output
	}
	inApp := response.Receipt.InApp[0]
	item := response.LatestReceiptInfo[0]
	defer tx.Rollback()
	// 驗證該用戶訂閱狀態
	findSubscribeInfoInput := subscribeInfoModel.FindInput{}
	findSubscribeInfoInput.UserID = util.PointerInt64(input.UserID)
	subscribeInfoOutput, err := r.subscribeInfoService.Tx(tx).Find(&findSubscribeInfoInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(subscribeInfoOutput.Status, 0) == 1 {
		output.Set(code.BadRequest, "該用戶為訂閱狀態")
		return output
	}
	// 驗證被 OriginalTransactionID 綁定的用戶是否還在訂閱狀態
	subscribeInfoListInput := subscribeInfoModel.ListInput{}
	subscribeInfoListInput.OriginalTransactionID = util.PointerString(inApp.OriginalTransactionID)
	subscribeInfoOutputs, _, err := r.subscribeInfoService.Tx(tx).List(&subscribeInfoListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, subscribeInfoOutput := range subscribeInfoOutputs {
		if util.OnNilJustReturnInt(subscribeInfoOutput.Status, 0) == 1 {
			output.Set(code.BadRequest, "該 original_transaction_id 為訂閱狀態")
			return output
		}
	}
	// 驗證是否為訂閱收據
	subscribePlanListInput := subscribePlanModel.ListInput{}
	subscribePlanListInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN product_labels ON subscribe_plans.product_label_id = product_labels.id"},
	}
	subscribePlanListInput.Wheres = []*whereModel.Where{
		{Query: "product_labels.product_id = ?", Args: []interface{}{inApp.ProductID}},
	}
	subscribePlanOutputs, _, err := r.subscribePlanService.Tx(tx).List(&subscribePlanListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(subscribePlanOutputs) == 0 {
		output.Set(code.BadRequest, "訂單類型錯誤")
	}
	// 驗證收據是否已被使用
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.OriginalTransactionID = util.PointerString(inApp.OriginalTransactionID)
	receiptListInput.TransactionID = util.PointerString(inApp.TransactionID)
	receiptOutputs, _, err := r.receiptService.Tx(tx).List(&receiptListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(receiptOutputs) > 0 {
		output.Set(code.BadRequest, "此收據已有上傳記錄")
		return output
	}
	// 0.清空原先綁定該 OriginalTransactionID 的用戶
	subscribeInfoUpdateTables := make([]*subscribeInfoModel.Table, 0)
	for _, subscribeInfoOutput := range subscribeInfoOutputs {
		subscribeInfoUpdateTable := subscribeInfoModel.Table{}
		subscribeInfoUpdateTable.UserID = subscribeInfoOutput.UserID
		subscribeInfoUpdateTable.OriginalTransactionID = util.PointerString("")
		subscribeInfoUpdateTables = append(subscribeInfoUpdateTables, &subscribeInfoUpdateTable)
	}
	if err := r.subscribeInfoService.Tx(tx).Updates(subscribeInfoUpdateTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 1.創建訂單
	orderID := time.Now().Format("20060102150405") + strconv.Itoa(int(util.RandRange(100000, 999999)))
	table := orderModel.Table{}
	table.ID = util.PointerString(orderID)
	table.UserID = util.PointerInt64(input.UserID)
	table.Quantity = util.PointerInt(1)
	table.Type = util.PointerInt(orderModel.Subscribe)
	table.OrderStatus = util.PointerInt(orderModel.Success)
	_, err = r.orderService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.創建訂單與訂閱項目關聯
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = util.PointerString(orderID)
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutputs[0].ID
	orderSubscribePlanTable.Status = util.PointerInt(1)
	if err := r.orderSubscribePlanService.Tx(tx).Create(&orderSubscribePlanTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3.儲存收據
	quantity, _ := strconv.Atoi(item.Quantity)
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = util.PointerString(orderID)
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAP)
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	receiptTable.TransactionID = util.PointerString(item.TransactionID)
	receiptTable.ProductID = util.PointerString(item.ProductID)
	receiptTable.Quantity = util.PointerInt(quantity)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4.更新用戶訂閱狀態與OriginalTransactionID
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = util.PointerInt64(input.UserID)
	subscribeInfoTable.OrderID = util.PointerString(orderID)
	subscribeInfoTable.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.ValidSubscribe)
	subscribeInfoTable.PaymentType = util.PointerInt(receiptModel.IAP)
	subscribeInfoTable.StartDate = util.PointerString(item.PurchaseDate.Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(item.ExpiresDate.Format("2006-01-02 15:04:05"))
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5.更新用戶類型
	userTable := userModel.Table{}
	userTable.ID = util.PointerInt64(input.UserID)
	userTable.UserType = util.PointerInt(userModel.Subscribe)
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.BadRequest, "success!")
	return output
}

func (r *resolver) APIUploadAppleChargeReceipt(ctx *gin.Context, tx *gorm.DB, input *api_upload_apple_charge_receipt.Input) (output api_upload_apple_charge_receipt.Output) {
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
	if util.OnNilJustReturnInt(orderOutput.Type, 0) != orderModel.BuyCourse {
		output.Set(code.BadRequest, "此訂單非訂閱訂單，訂單類型錯誤")
		return output
	}
	if err := r.handleBuyCourseTradeForApple(tx, orderOutput, response); err != nil {
		logger.Shared().Error(ctx, "APIVerifyAppleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIVerifyGoogleReceipt(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIVerifyGoogleReceiptInput) (output orderModel.APIVerifyGoogleReceiptOutput) {
	defer tx.Rollback()
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
		logger.Shared().Error(ctx, "APIVerifyGoogleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取API Token
	token, err := r.iabTool.APIGetGooglePlayToken(oauthToken)
	if err != nil {
		logger.Shared().Error(ctx, "APIVerifyGoogleReceipt："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) == orderModel.BuyCourse {
		response, err := r.iabTool.APIGetProducts(input.Body.ProductID, input.Body.ReceiptData, token)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
		}
		if err := r.handleBuyCourseTradeForGoogle(tx, orderOutput, response); err != nil {
			logger.Shared().Error(ctx, "APIVerifyGoogleReceipt："+err.Error())
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
			return output
		}
		if err := r.handleSubscribeTradeForGoogle(tx, input.Body.ProductID, orderOutput, response, input.Body.ReceiptData); err != nil {
			logger.Shared().Error(ctx, "APIVerifyGoogleReceipt："+err.Error())
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

func (r *resolver) APIOrderRedeem(tx *gorm.DB, input *api_order_redeem.Input) (output api_order_redeem.Output) {
	// 查詢訂單
	findInput := orderModel.FindInput{}
	findInput.ID = util.PointerString(input.Uri.OrderID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderCourse.Course"},
		{Field: "OrderCourse.SaleItem.ProductLabel"},
	}
	orderOutput, err := r.orderService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證訂單
	if util.OnNilJustReturnInt64(orderOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非該用戶訂單")
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.Type, 0) != orderModel.BuyCourse {
		output.Set(code.BadRequest, "此訂單類型錯誤，無法兌換課表")
		return output
	}
	if orderOutput.OrderCourse == nil {
		output.Set(code.BadRequest, "此訂單的未註名商品")
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.OrderCourseOnSafe().CourseOnSafe().SaleType, 0) != courseModel.SaleTypeFree {
		output.Set(code.BadRequest, "此訂單的商品非免費課表")
		return output
	}
	if util.OnNilJustReturnInt(orderOutput.OrderStatus, 0) != orderModel.Pending {
		output.Set(code.BadRequest, "此訂單已失效")
		return output
	}
	defer tx.Rollback()
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = orderOutput.ID
	receiptTable.PaymentType = util.PointerInt(0)
	receiptTable.OriginalTransactionID = util.PointerString("")
	receiptTable.TransactionID = util.PointerString("")
	receiptTable.ReceiptToken = util.PointerString("")
	receiptTable.ProductID = util.PointerString("")
	receiptTable.Quantity = util.PointerInt(1)
	_, err = r.receiptService.Tx(tx).CreateOrUpdate(&receiptTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.創建購買紀錄
	logTable := purchaseLogModel.Table{}
	logTable.UserID = orderOutput.UserID
	logTable.OrderID = orderOutput.ID
	logTable.Type = util.PointerInt(purchaseLogModel.Buy)
	_, err = r.purchaseLogService.Tx(tx).Create(&logTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3.創建購買資源
	assetTable := courseAssetModel.Table{}
	assetTable.UserID = orderOutput.UserID
	assetTable.CourseID = orderOutput.OrderCourseOnSafe().CourseID
	assetTable.Available = util.PointerInt(1)
	_, err = r.courseAssetService.Tx(tx).Create(&assetTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = orderOutput.ID
	orderTable.OrderStatus = util.PointerInt(orderModel.Success)
	if err := r.orderService.Tx(tx).Update(&orderTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIAppStoreNotification(ctx *gin.Context, tx *gorm.DB, input *orderModel.APIAppStoreNotificationInput) (output orderModel.APIAppStoreNotificationOutput) {
	defer tx.Rollback()
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
	// 查詢當前綁定 OriginalTransactionId 的用戶
	subscribeInfoListInput := subscribeInfoModel.ListInput{}
	subscribeInfoListInput.OriginalTransactionID = util.PointerString(response.Data.SignedTransactionInfo.OriginalTransactionId)
	subscribeInfoListInput.Size = 1
	subscribeInfoListInput.Page = 1
	subscribeInfoListInput.OrderField = "update_at"
	subscribeInfoListInput.OrderType = order_by.DESC
	subscribeInfoOutputs, _, err := r.subscribeInfoService.List(&subscribeInfoListInput)
	if err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(subscribeInfoOutputs) == 0 {
		output.Set(code.BadRequest, "當前無綁定此 OriginalTransactionId 用戶")
		return output
	}
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = subscribeInfoOutputs[0].OrderID
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
	// 2.修改訂單訂閱項目(升級or降級, 訂閱狀態)
	var subscribePlanStatus = util.PointerInt(1)
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		subscribePlanStatus = util.PointerInt(0)
	}
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = subscribeInfoOutputs[0].OrderID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	orderSubscribePlanTable.Status = subscribePlanStatus
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		logger.Shared().Error(ctx, "APIGooglePlayNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3.更新用戶訂閱狀態
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = subscribeInfoOutputs[0].UserID
	subscribeInfoTable.OrderID = subscribeInfoOutputs[0].OrderID
	subscribeInfoTable.OriginalTransactionID = util.PointerString(response.Data.SignedTransactionInfo.OriginalTransactionId)
	subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.ValidSubscribe)
	subscribeInfoTable.PaymentType = util.PointerInt(receiptModel.IAP)
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(response.Data.SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		subscribeInfoTable.OriginalTransactionID = util.PointerString("")
		subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.NoneSubscribe)
	}
	if err := r.subscribeInfoService.Tx(tx).CreateOrUpdate(&subscribeInfoTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4.更新用戶類型
	userTable := userModel.Table{}
	userTable.ID = subscribeInfoOutputs[0].UserID
	userTable.UserType = util.PointerInt(userModel.Subscribe)
	if subscribeLogType == subscribeLogModel.Expired || subscribeLogType == subscribeLogModel.Refund {
		userTable.UserType = util.PointerInt(userModel.Normal)
	}
	if err := r.userService.Tx(tx).Update(&userTable); err != nil {
		logger.Shared().Error(ctx, "APIAppStoreNotification："+err.Error())
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5.更新訂單狀態
	orderTable := orderModel.Table{}
	orderTable.ID = subscribeInfoOutputs[0].OrderID
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
	defer tx.Rollback()
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
	subscribeLogTable.PurchaseDate = util.PointerString(util.UnixToTime(startTimeMillis / 1000).Format("2006-01-02 15:04:05"))
	subscribeLogTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis / 1000).Format("2006-01-02 15:04:05"))
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
	subscribeInfoTable.PaymentType = util.PointerInt(receiptModel.IAB)
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(startTimeMillis / 1000).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis / 1000).Format("2006-01-02 15:04:05"))
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
			findUserInput := userModel.FindInput{}
			findUserInput.ID = orderOutputs[0].UserID
			userOutput, err := r.userService.Find(&findUserInput)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			var accountType string
			switch util.OnNilJustReturnInt(userOutput.AccountType, 0) {
			case userModel.Email:
				accountType = "Email"
			case userModel.Facebook:
				accountType = "Facebook"
			case userModel.Google:
				accountType = "Google"
			case userModel.Line:
				accountType = "Line"
			case userModel.Apple:
				accountType = "Apple"
			}
			msg := fmt.Sprintf("此 Apple ID 已綁定 %v 信箱( %v 註冊)", util.OnNilJustReturnString(userOutput.Email, ""), accountType)
			output.Set(code.BadRequest, msg)
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

func (r *resolver) SyncAppleSubscribeStatusSchedule(tx *gorm.DB) {
	defer tx.Rollback()
	subscribeInfoList := subscribeInfoModel.ListInput{}
	subscribeInfoList.Status = util.PointerInt(1)      // 會員狀態(0:無會員狀態/1:付費會員狀態)
	subscribeInfoList.PaymentType = util.PointerInt(1) // 支付方式(0:尚未指定/1:apple內購/2:google內購)
	subscribeInfoList.Wheres = []*whereModel.Where{
		{Query: "user_subscribe_infos.expires_date <= ?", Args: []interface{}{time.Now().Format("2006-01-02 15:04:05")}},
	}
	subscribeInfoOutputs, _, err := r.subscribeInfoService.Tx(tx).List(&subscribeInfoList)
	if err != nil {
		return
	}
	token, err := r.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return
	}
	for _, subscribeInfoOutput := range subscribeInfoOutputs {
		// 查詢 OriginalTransactionID 訂閱狀態
		response, _ := r.iapTool.GetSubscribeAPI(util.OnNilJustReturnString(subscribeInfoOutput.OriginalTransactionID, ""), token)
		// 準備 user_subscribe_info table
		subscribeInfoTable := subscribeInfoModel.Table{}
		subscribeInfoTable.UserID = subscribeInfoOutput.UserID
		subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.NoneSubscribe)
		subscribeInfoTable.OriginalTransactionID = util.PointerString("")
		// 準備 user table
		userTable := userModel.Table{}
		userTable.ID = subscribeInfoOutput.UserID
		userTable.UserType = util.PointerInt(userModel.Normal)
		if response != nil {
			if len(response.Data) > 0 {
				if len(response.Data[0].LastTransactions) > 0 {
					status := response.Data[0].LastTransactions[0].Status
					subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(response.Data[0].LastTransactions[0].SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
					subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(response.Data[0].LastTransactions[0].SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
					if status == 1 || status == 3 || status == 4 || status == 5 { // 當前訂閱尚未過期
						subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.ValidSubscribe)
						subscribeInfoTable.OriginalTransactionID = subscribeInfoOutput.OriginalTransactionID
						userTable.UserType = util.PointerInt(userModel.Subscribe)
					}
				}
			}
		}
		// 更新 user_subscribe_info
		if err := r.subscribeInfoService.Tx(tx).Update(&subscribeInfoTable); err != nil {
			return
		}
		// 更新用戶類型
		if err := r.userService.Tx(tx).Update(&userTable); err != nil {
			return
		}
	}
	tx.Commit()
}

func (r *resolver) handleBuyCourseTradeForApple(tx *gorm.DB, order *orderModel.Output, response *iapModel.IAPVerifyReceiptResponse) error {
	defer tx.Rollback()
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
	defer tx.Rollback()
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
	defer tx.Rollback()
	// 驗證收據格式
	if len(response.LatestReceiptInfo) == 0 || len(response.PendingRenewalInfo) == 0 {
		return errors.New("無效的收據(無LatestReceiptInfo或PendingRenewalInfo參數)")
	}
	item := response.LatestReceiptInfo[0]
	if response.LatestReceiptInfo[0].ExpiresDate == nil {
		return errors.New("無效的收據(無ExpiresDate參數)")
	}
	if len(response.PendingRenewalInfo[0].ExpirationIntent) > 0 {
		return errors.New("訂閱已過期")
	}
	// 驗證被 OriginalTransactionID 綁定的用戶是否還在訂閱中狀態
	subscribeInfoListInput := subscribeInfoModel.ListInput{}
	subscribeInfoListInput.UserID = order.UserID
	subscribeInfoListInput.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	subscribeInfoOutputs, _, err := r.subscribeInfoService.Tx(tx).List(&subscribeInfoListInput)
	if err != nil {
		return err
	}
	for _, subscribeInfoOutput := range subscribeInfoOutputs {
		if util.OnNilJustReturnInt(subscribeInfoOutput.Status, 0) == 1 {
			return errors.New("該用戶已是訂閱狀態")
		}
	}
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
	findPlanInput := subscribePlanModel.FindInput{}
	findPlanInput.ProductID = util.PointerString(item.ProductID)
	subscribePlanOutput, err := r.subscribePlanService.Tx(tx).Find(&findPlanInput)
	if err != nil {
		return err
	}
	orderSubscribePlanTable := orderSubscribePlanModel.Table{}
	orderSubscribePlanTable.OrderID = order.ID
	orderSubscribePlanTable.SubscribePlanID = subscribePlanOutput.ID
	if err := r.orderSubscribePlanService.Tx(tx).Update(&orderSubscribePlanTable); err != nil {
		return err
	}
	// 清空原先綁定該 OriginalTransactionID 的用戶
	subscribeInfoUpdateTables := make([]*subscribeInfoModel.Table, 0)
	for _, subscribeInfoOutput := range subscribeInfoOutputs {
		subscribeInfoUpdateTable := subscribeInfoModel.Table{}
		subscribeInfoUpdateTable.UserID = subscribeInfoOutput.UserID
		subscribeInfoUpdateTable.OriginalTransactionID = util.PointerString("")
		subscribeInfoUpdateTables = append(subscribeInfoUpdateTables, &subscribeInfoUpdateTable)
	}
	if err := r.subscribeInfoService.Tx(tx).Updates(subscribeInfoUpdateTables); err != nil {
		return err
	}
	// 3.更新用戶訂閱狀態與綁定狀態
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.OriginalTransactionID = util.PointerString("")
	subscribeInfoTable.UserID = order.UserID
	subscribeInfoTable.OrderID = order.ID
	subscribeInfoTable.OriginalTransactionID = util.PointerString(item.OriginalTransactionID)
	subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.ValidSubscribe)
	subscribeInfoTable.PaymentType = util.PointerInt(receiptModel.IAP)
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
	// 6.更新訂單訂閱狀態
	subscribePlanTable := orderSubscribePlanModel.Table{}
	subscribePlanTable.OrderID = order.ID
	subscribePlanTable.Status = util.PointerInt(subscribeInfoModel.ValidSubscribe)
	if err := r.orderSubscribePlanService.Tx(tx).Update(&subscribePlanTable); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (r *resolver) handleSubscribeTradeForGoogle(tx *gorm.DB, productID string, order *orderModel.Output, response *iabModel.IABSubscriptionAPIResponse, receiptToken string) error {
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
	// 驗證是否是原先訂閱用戶
	orderListInput := orderModel.ListInput{}
	orderListInput.OriginalTransactionID = util.PointerString(originalTransactionID)
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
	findPlanInput.ProductID = util.PointerString(productID)
	subscribePlanOutput, err := r.subscribePlanService.Find(&findPlanInput)
	if err != nil {
		return err
	}
	// 1.儲存收據
	receiptTable := receiptModel.Table{}
	receiptTable.OrderID = order.ID
	receiptTable.PaymentType = util.PointerInt(receiptModel.IAB)
	receiptTable.OriginalTransactionID = util.PointerString(originalTransactionID)
	receiptTable.TransactionID = util.PointerString(transactionID)
	receiptTable.ReceiptToken = util.PointerString(receiptToken)
	receiptTable.ProductID = util.PointerString(productID)
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
	var subscribeStatus = subscribeInfoModel.NoneSubscribe
	if response.PaymentState == 1 || response.PaymentState == 2 {
		subscribeStatus = subscribeInfoModel.ValidSubscribe
	}
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = order.UserID
	subscribeInfoTable.OrderID = order.ID
	subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
	subscribeInfoTable.PaymentType = util.PointerInt(receiptModel.IAB)
	startTimeMillis, err := strconv.ParseInt(response.StartTimeMillis, 10, 64)
	if err != nil {
		return nil
	}
	expiryTimeMillis, err := strconv.ParseInt(response.ExpiryTimeMillis, 10, 64)
	if err != nil {
		return nil
	}
	subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(startTimeMillis / 1000).Format("2006-01-02 15:04:05"))
	subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis / 1000).Format("2006-01-02 15:04:05"))
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
