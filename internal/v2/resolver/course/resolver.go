package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	planModel "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	saleItemModel "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	workoutModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type resolver struct {
	courseService        courseService.Service
	planService          plan.Service
	workoutService       workout.Service
	subscribeInfoService user_subscribe_info.Service
	saleItemService      sale_item.Service
	uploadTool           uploader.Tool
}

func New(courseService courseService.Service, planService plan.Service,
	workoutService workout.Service, subscribeInfoService user_subscribe_info.Service,
	saleItemService sale_item.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{courseService: courseService, planService: planService,
		workoutService: workoutService, subscribeInfoService: subscribeInfoService,
		saleItemService: saleItemService, uploadTool: uploadTool}
}

func (r *resolver) APIGetFavoriteCourses(input *model.APIGetFavoriteCoursesInput) (output model.APIGetFavoriteCoursesOutput) {
	// parser input
	param := model.FavoriteListInput{}
	param.UserID = util.PointerInt64(input.UserID)
	param.OrderField = "create_at"
	param.OrderType = order_by.DESC
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 執行查詢
	results, page, err := r.courseService.FavoriteList(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetFavoriteCoursesData{}
	if err := util.Parser(results, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetCMSCourses(ctx *gin.Context, input *model.APIGetCMSCoursesInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.IgnoredCourseStatus = []int{1}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, page, err := r.courseService.List(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCoursesData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCoursesOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}

func (r *resolver) APIGetCMSCourse(ctx *gin.Context, input *model.APIGetCMSCourseInput) interface{} {
	param := model.FindInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, err := r.courseService.Find(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCourseData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCourseOutput{}
	output.Data = &data
	output.Code = code.Success
	output.Msg = "success!"
	return output
}

func (r *resolver) APIUpdateCMSCoursesStatus(input *model.APIUpdateCMSCoursesStatusInput) (output base.Output) {
	tables := make([]*model.Table, 0)
	for _, courseID := range input.IDs {
		table := model.Table{}
		table.ID = util.PointerInt64(courseID)
		table.CourseStatus = &input.CourseStatus
		tables = append(tables, &table)
	}
	if err := r.courseService.Updates(tables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIUpdateCMSCourseCover(input *model.APIUpdateCMSCourseCoverInput) (output model.APIUpdateCMSCourseCoverOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Cover = util.PointerString(fileNamed)
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}

func (r *resolver) APICreateUserCourse(input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput) {
	/** 驗證是否能創建 */
	// 1. 檢查目前創建多計畫課表數量
	courseListInput := model.ListInput{}
	courseListInput.UserID = util.PointerInt64(input.UserID)
	courseListInput.SaleType = util.PointerInt(model.SaleTypePersonal)
	courseListInput.ScheduleType = util.PointerInt(2)
	courseOutputs, _, err := r.courseService.List(&courseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 檢查目前是否已訂閱
	subscribeListInput := subscribeInfoModel.ListInput{}
	subscribeListInput.UserID = util.PointerInt64(input.UserID)
	subscribeListInput.Page = 1
	subscribeListInput.Size = 1
	subscribeListOutput, _, err := r.subscribeInfoService.List(&subscribeListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3. 驗證多計畫課表創建是否達上線
	if len(subscribeListOutput) == 0 && len(courseOutputs) >= 1 {
		output.Set(code.PermissionDenied, "多計畫課表創建已達上限(未訂閱)")
		return output
	}
	if util.OnNilJustReturnInt(subscribeListOutput[0].Status, 0) == 0 && len(courseOutputs) >= 1 {
		output.Set(code.PermissionDenied, "多計畫課表創建已達上限(未訂閱)")
		return output
	}
	// 創建課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypePersonal)
	table.ScheduleType = util.PointerInt(model.MultiplePlan)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(0)
	table.WorkoutCount = util.PointerInt(0)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateUserCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateUserSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput) {
	defer tx.Rollback()
	//創建單一訓練課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypePersonal)
	table.ScheduleType = util.PointerInt(model.SingleWorkout)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(1)
	table.WorkoutCount = util.PointerInt(1)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一計畫
	planTable := planModel.Table{}
	planTable.Name = util.PointerString("單一計畫")
	planTable.CourseID = util.PointerInt64(courseID)
	planTable.WorkoutCount = util.PointerInt(1)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一訓練
	workoutTable := workoutModel.Table{}
	workoutTable.PlanID = util.PointerInt64(planID)
	workoutTable.Name = util.PointerString("單一訓練")
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	_, err = r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	//Parser Output
	data := model.APICreateUserCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetUserPersonalCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢個人課表
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.SaleType = util.PointerInt(model.SaleTypePersonal)
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetUserProgressCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢進行中課表
	listInput := model.ProgressListInput{}
	listInput.UserID = input.UserID
	listInput.OrderField = "update_at"
	listInput.OrderType = order_by.DESC
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
		{Field: "SaleItem.ProductLabel"},
	}
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.ProgressList(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetUserChargeCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢付費課表
	listInput := model.ChargeListInput{}
	listInput.UserID = input.UserID
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
		{Field: "SaleItem.ProductLabel"},
	}
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.ChargeList(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIDeleteUserCourse(input *model.APIDeleteUserCourseInput) (output model.APIDeleteUserCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法刪除資源")
		return output
	}
	// 刪除課表
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.courseService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIUpdateUserCourse(input *model.APIUpdateUserCourseInput) (output model.APIUpdateUserCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法刪除資源")
		return output
	}
	// 修改課表
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetUserCourse(input *model.APIGetUserCourseInput) (output model.APIGetUserCourseOutput) {
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "UserCourseStatistic", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCourseData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetTrainerCourses(input *model.APIGetTrainerCoursesInput) (output model.APIGetTrainerCoursesOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.CourseStatus = input.Query.CourseStatus
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetTrainerCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APICreateTrainerCourse(input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput) {
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypeNone)
	table.ScheduleType = util.PointerInt(model.MultiplePlan)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(0)
	table.WorkoutCount = util.PointerInt(0)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateTrainerCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateTrainerSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput) {
	defer tx.Rollback()
	//創建單一訓練課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypeNone)
	table.ScheduleType = util.PointerInt(model.SingleWorkout)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(1)
	table.WorkoutCount = util.PointerInt(1)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一計畫
	planTable := planModel.Table{}
	planTable.Name = util.PointerString("單一計畫")
	planTable.CourseID = util.PointerInt64(courseID)
	planTable.WorkoutCount = util.PointerInt(1)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一訓練
	workoutTable := workoutModel.Table{}
	workoutTable.PlanID = util.PointerInt64(planID)
	workoutTable.Name = util.PointerString("單一訓練")
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	_, err = r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	//Parser Output
	data := model.APICreateTrainerCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetTrainerCourse(input *model.APIGetTrainerCourseInput) (output model.APIGetTrainerCourseOutput) {
	// 獲取課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證
	if input.UserID != util.OnNilJustReturnInt64(courseOutput.UserID, 0) {
		output.Set(code.BadRequest, "非該課表擁有者，無法查看課表資訊")
		return output
	}
	// parser output
	data := model.APIGetTrainerCourseData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateTrainerCourse(tx *gorm.DB, input *model.APIUpdateTrainerCourseInput) (output model.APIUpdateTrainerCourseOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := model.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt(input.Form.SaleType, 0) == model.SaleTypePersonal {
		output.Set(code.BadRequest, "教練課表無法修改為個人課表")
		return output
	}
	// 驗證輸入參數
	if util.OnNilJustReturnInt(input.Form.SaleType, 0) == model.SaleTypeCharge {
		if input.Form.SaleID == nil {
			output.Set(code.BadRequest, "當銷售類型為付費課表時，SaleID必須帶值")
			return output
		}
		findSaleItemInput := saleItemModel.FindInput{}
		findSaleItemInput.ID = input.Form.SaleID
		saleItemOutput, err := r.saleItemService.Find(&findSaleItemInput)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		if util.OnNilJustReturnInt(saleItemOutput.Type, 0) != model.SaleTypeCharge {
			output.Set(code.BadRequest, "銷售類型不符")
			return output
		}
	}
	// 修改課表
	table := model.Table{}
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.courseService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//當課表類型不為付費課表時，sale_id 保持 nil
	findCourseInput = model.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err = r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypeCharge {
		if err := r.courseService.Tx(tx).UpdateSaleID(input.Uri.ID, nil); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	// 上傳 cover
	if input.Cover != nil {
		// 儲存新 cover
		newCoverNamed, err := r.uploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改課表
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.Cover = util.PointerString(newCoverNamed)
		if err := r.courseService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 cover
		_ = r.uploadTool.Delete(util.OnNilJustReturnString(courseOutput.Cover, ""))
	}
	tx.Commit()
	// parser output
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem.ProductLabel"},
	}
	courseOutput, err = r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIUpdateTrainerCourseData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Data = &data
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteTrainerCourse(input *model.APIDeleteTrainerCourseInput) (output model.APIDeleteTrainerCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 刪除課表
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.courseService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APISubmitTrainerCourse(input *model.APISubmitTrainerCourseInput) (output model.APISubmitTrainerCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 檢查課表狀態
	saleType := util.OnNilJustReturnInt(courseOutput.SaleType, 0)
	if saleType == model.SaleTypeNone {
		output.Set(code.BadRequest, "未設定 sale_type")
		return output
	}
	if saleType == model.SaleTypeCharge && courseOutput.SaleID == nil {
		output.Set(code.BadRequest, "未設定 sale_item_id")
		return output
	}
	courseStatus := util.OnNilJustReturnInt(courseOutput.CourseStatus, 0)
	if courseStatus != model.Preparing && courseStatus != model.Reject {
		output.Set(code.BadRequest, "此課表無法重新送審")
		return output
	}
	// 更新課表
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.CourseStatus = util.PointerInt(model.Reviewing)
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
