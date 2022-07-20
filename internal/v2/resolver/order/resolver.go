package order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	orderModel "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	orderCourseModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	courseAssetModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type resolver struct {
	orderService order.Service
	courseService courseService.Service
	orderCourserService order_course.Service
	courseAssetService user_course_asset.Service
}

func New(orderService order.Service, courseService courseService.Service, orderCourserService order_course.Service, courseAssetService user_course_asset.Service) Resolver {
	return &resolver{orderService: orderService, courseService: courseService, orderCourserService: orderCourserService, courseAssetService: courseAssetService}
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
	if util.OnNilJustReturnInt(courseOutput.CourseStatus,0)  != courseModel.Sale {
		output.Set(code.BadRequest, "必須為銷售狀態的課表才可被加入訂單")
		return output
	}
	if  util.OnNilJustReturnInt(courseOutput.SaleType,0) != courseModel.SaleTypeFree && util.OnNilJustReturnInt(courseOutput.SaleType,0) != courseModel.SaleTypeCharge {
		output.Set(code.BadRequest, "商品必須為免費課表或付費課表類型才可創建此訂單")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType,0) != courseModel.SaleTypeFree && courseOutput.SaleID == nil {
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
