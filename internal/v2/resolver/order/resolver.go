package order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	orderModel "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderCourseModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	orderSubscribePlanModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	courseAssetModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type resolver struct {
	orderService              order.Service
	courseService             course.Service
	orderCourserService       order_course.Service
	courseAssetService        user_course_asset.Service
	subscribeInfoService      user_subscribe_info.Service
	orderSubscribePlanService order_subscribe_plan.Service
}

func New(orderService order.Service, courseService course.Service,
	orderCourserService order_course.Service, courseAssetService user_course_asset.Service,
	subscribeInfoService user_subscribe_info.Service, orderSubscribePlanService order_subscribe_plan.Service) Resolver {
	return &resolver{orderService: orderService, courseService: courseService,
		orderCourserService: orderCourserService, courseAssetService: courseAssetService,
		subscribeInfoService: subscribeInfoService, orderSubscribePlanService: orderSubscribePlanService}
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
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypeFree && courseOutput.SaleID == nil {
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
	// 檢查是否有創建過訂閱訂單，有則直接回傳(一個用戶最多只會產生一個訂閱的訂單)
	orderListInput := orderModel.ListInput{}
	orderListInput.UserID = util.PointerInt64(input.UserID)
	orderListInput.Type = util.PointerInt(orderModel.Subscribe)
	orderListInput.OrderField = "create_at"
	orderListInput.OrderType = order_by.DESC
	orderListInput.Page = 1
	orderListInput.Size = 1
	orderListInput.Preloads = []*preloadModel.Preload{
		{Field: "OrderSubscribePlan.SubscribePlan.ProductLabel"},
	}
	orderOutputs, _, err := r.orderService.List(&orderListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(orderOutputs) > 0 {
		data := orderModel.APICreateSubscribeOrderData{}
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
